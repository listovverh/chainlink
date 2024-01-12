package chainlink

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pelletier/go-toml/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos"
	coscfg "github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/config"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana"
	pkgsolana "github.com/smartcontractkit/chainlink-solana/pkg/solana"
	pkgstarknet "github.com/smartcontractkit/chainlink-starknet/relayer/pkg/chainlink"
	starkchain "github.com/smartcontractkit/chainlink-starknet/relayer/pkg/chainlink/chain"
	"github.com/smartcontractkit/chainlink-starknet/relayer/pkg/chainlink/config"

	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/config/env"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
	evmrelay "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/wsrpc"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type RelayerFactory struct {
	logger.Logger
	*plugins.LoopRegistry
	loop.GRPCOpts
	MercuryPool wsrpc.Pool
}

type EVMFactoryConfig struct {
	legacyevm.ChainOpts
	evmrelay.CSAETHKeystore
}

func (r *RelayerFactory) NewEVM(ctx context.Context, config EVMFactoryConfig) (map[relay.ID]evmrelay.LOOPRelayAdapter, error) {
	// TODO impl EVM loop. For now always 'fallback' to an adapter and embedded chain

	relayers := make(map[relay.ID]evmrelay.LOOPRelayAdapter)

	lggr := r.Logger.Named("EVM")

	// override some common opts with the factory values. this seems weird... maybe other signatures should change, or this should take a different type...
	ccOpts := legacyevm.ChainRelayOpts{
		Logger:    lggr,
		KeyStore:  config.CSAETHKeystore.Eth(),
		ChainOpts: config.ChainOpts,
	}

	legacyChains, err := evmrelay.NewLegacyChains(ctx, ccOpts)
	if err != nil {
		return nil, err
	}
	for _, chain := range legacyChains {
		relayID := relay.ID{Network: relay.EVM, ChainID: chain.ID().String()}
		chain := chain

		relayerOpts := evmrelay.RelayerOpts{
			DB:             ccOpts.SqlxDB,
			QConfig:        ccOpts.AppConfig.Database(),
			CSAETHKeystore: config.CSAETHKeystore,
			MercuryPool:    r.MercuryPool,
		}
		relayer, err2 := evmrelay.NewRelayer(lggr.Named(relayID.ChainID), chain, relayerOpts)
		if err2 != nil {
			err = errors.Join(err, err2)
			continue
		}

		relayers[relayID] = evmrelay.NewLOOPRelayAdapter(relayer)
	}

	// always return err because it is accumulating individual errors
	return relayers, err
}

type SolanaFactoryConfig struct {
	Keystore keystore.Solana
	solana.TOMLConfigs
}

func (r *RelayerFactory) NewSolana(ks keystore.Solana, chainCfgs solana.TOMLConfigs) (map[relay.ID]loop.Relayer, error) {
	solanaRelayers := make(map[relay.ID]loop.Relayer)
	var (
		solLggr = r.Logger.Named("Solana")
		signer  = &keystore.SolanaSigner{Solana: ks}
	)

	unique := make(map[string]struct{})
	// create one relayer per chain id
	for _, chainCfg := range chainCfgs {

		relayID := relay.ID{Network: relay.Solana, ChainID: *chainCfg.ChainID}
		_, alreadyExists := unique[relayID.Name()]
		if alreadyExists {
			return nil, fmt.Errorf("duplicate chain definitions for %s", relayID.Name())
		}
		unique[relayID.Name()] = struct{}{}

		// skip disabled chains from further processing
		if !chainCfg.IsEnabled() {
			solLggr.Warnw("Skipping disabled chain", "id", chainCfg.ChainID)
			continue
		}

		lggr := solLggr.Named(relayID.ChainID)

		if cmdName := env.SolanaPlugin.Cmd.Get(); cmdName != "" {

			// setup the solana relayer to be a LOOP
			cfgTOML, err := toml.Marshal(struct {
				Solana solana.TOMLConfig
			}{Solana: *chainCfg})

			if err != nil {
				return nil, fmt.Errorf("failed to marshal Solana configs: %w", err)
			}
			envVars, err := plugins.ParseEnvFile(env.SolanaPlugin.Env.Get())
			if err != nil {
				return nil, fmt.Errorf("failed to parse Solana env file: %w", err)
			}
			solCmdFn, err := plugins.NewCmdFactory(r.Register, plugins.CmdConfig{
				ID:  relayID.Name(),
				Cmd: cmdName,
				Env: envVars,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create Solana LOOP command: %w", err)
			}

			solanaRelayers[relayID] = loop.NewRelayerService(lggr, r.GRPCOpts, solCmdFn, string(cfgTOML), signer)

		} else {
			// fallback to embedded chain
			opts := solana.ChainOpts{
				Logger:   lggr,
				KeyStore: signer,
			}

			chain, err := solana.NewChain(chainCfg, opts)
			if err != nil {
				return nil, err
			}
			solanaRelayers[relayID] = relay.NewServerAdapter(pkgsolana.NewRelayer(lggr, chain))
		}
	}
	return solanaRelayers, nil
}

type StarkNetFactoryConfig struct {
	Keystore keystore.StarkNet
	config.TOMLConfigs
}

// TODO BCF-2606 consider consolidating the driving logic with that of NewSolana above via generics
// perhaps when we implement a Cosmos LOOP
func (r *RelayerFactory) NewStarkNet(ks keystore.StarkNet, chainCfgs config.TOMLConfigs) (map[relay.ID]loop.Relayer, error) {
	starknetRelayers := make(map[relay.ID]loop.Relayer)

	var (
		starkLggr = r.Logger.Named("StarkNet")
		loopKs    = &keystore.StarknetLooppSigner{StarkNet: ks}
	)

	unique := make(map[string]struct{})
	// create one relayer per chain id
	for _, chainCfg := range chainCfgs {
		relayID := relay.ID{Network: relay.StarkNet, ChainID: *chainCfg.ChainID}
		_, alreadyExists := unique[relayID.Name()]
		if alreadyExists {
			return nil, fmt.Errorf("duplicate chain definitions for %s", relayID.Name())
		}
		unique[relayID.Name()] = struct{}{}

		// skip disabled chains from further processing
		if !chainCfg.IsEnabled() {
			starkLggr.Warnw("Skipping disabled chain", "id", chainCfg.ChainID)
			continue
		}

		lggr := starkLggr.Named(relayID.ChainID)

		if cmdName := env.StarknetPlugin.Cmd.Get(); cmdName != "" {
			// setup the starknet relayer to be a LOOP
			cfgTOML, err := toml.Marshal(struct {
				Starknet config.TOMLConfig
			}{Starknet: *chainCfg})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal StarkNet configs: %w", err)
			}

			envVars, err := plugins.ParseEnvFile(env.StarknetPlugin.Env.Get())
			if err != nil {
				return nil, fmt.Errorf("failed to parse Starknet env file: %w", err)
			}
			starknetCmdFn, err := plugins.NewCmdFactory(r.Register, plugins.CmdConfig{
				ID:  relayID.Name(),
				Cmd: cmdName,
				Env: envVars,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create StarkNet LOOP command: %w", err)
			}
			// the starknet relayer service has a delicate keystore dependency. the value that is passed to NewRelayerService must
			// be compatible with instantiating a starknet transaction manager KeystoreAdapter within the LOOPp executable.
			starknetRelayers[relayID] = loop.NewRelayerService(lggr, r.GRPCOpts, starknetCmdFn, string(cfgTOML), loopKs)
		} else {
			// fallback to embedded chain
			opts := starkchain.ChainOpts{
				Logger:   lggr,
				KeyStore: loopKs,
			}

			chain, err := starkchain.NewChain(chainCfg, opts)
			if err != nil {
				return nil, err
			}

			starknetRelayers[relayID] = relay.NewServerAdapter(pkgstarknet.NewRelayer(lggr, chain))
		}
	}
	return starknetRelayers, nil

}

type CosmosFactoryConfig struct {
	Keystore keystore.Cosmos
	coscfg.TOMLConfigs
	*sqlx.DB
	pg.QConfig
}

func (c CosmosFactoryConfig) Validate() error {
	var err error
	if c.Keystore == nil {
		err = errors.Join(err, fmt.Errorf("nil Keystore"))
	}
	if len(c.TOMLConfigs) == 0 {
		err = errors.Join(err, fmt.Errorf("no CosmosConfigs provided"))
	}
	if c.DB == nil {
		err = errors.Join(err, fmt.Errorf("nil DB"))
	}
	if c.QConfig == nil {
		err = errors.Join(err, fmt.Errorf("nil QConfig"))
	}

	if err != nil {
		err = fmt.Errorf("invalid CosmosFactoryConfig: %w", err)
	}
	return err
}

func (r *RelayerFactory) NewCosmos(config CosmosFactoryConfig) (map[relay.ID]LOOPRelayAdapter, error) {
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("cannot create Cosmos relayer: %w", err)
	}
	relayers := make(map[relay.ID]LOOPRelayAdapter)

	var (
		cosmosLggr = r.Logger.Named("Cosmos")
		loopKs     = &keystore.CosmosLoopKeystore{Cosmos: config.Keystore}
	)

	// create one relayer per chain id
	for _, chainCfg := range config.TOMLConfigs {
		relayID := relay.ID{Network: relay.Cosmos, ChainID: *chainCfg.ChainID}

		lggr := cosmosLggr.Named(relayID.ChainID)

		opts := cosmos.ChainOpts{
			Logger:   lggr,
			DB:       config.DB,
			KeyStore: loopKs,
		}

		chain, err := cosmos.NewChain(chainCfg, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to load Cosmos chain %q: %w", relayID, err)
		}

		relayers[relayID] = NewCosmosLOOPRelayerChain(cosmos.NewRelayer(lggr, chain))

	}
	return relayers, nil

}
