package evm

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	commonservices "github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

type ChainReaderService interface {
	services.ServiceCtx
	commontypes.ChainReader
}

type chainReader struct {
	lggr     logger.Logger
	lp       logpoller.LogPoller
	codec    commontypes.RemoteCodec
	client   evmclient.Client
	bindings Bindings
	commonservices.StateMachine
}

// NewChainReaderService is a constructor for ChainReader, returns nil if there is any error
func NewChainReaderService(lggr logger.Logger, lp logpoller.LogPoller, b Bindings, chain legacyevm.Chain, config types.ChainReaderConfig) (ChainReaderService, error) {
	parsed := &parsedTypes{
		encoderDefs: map[string]*codecEntry{},
		decoderDefs: map[string]*codecEntry{},
	}

	if err := addTypes(config.ChainContractReaders, b, parsed); err != nil {
		return nil, err
	}

	c, err := parsed.toCodec(lggr)

	return &chainReader{
		lggr:     lggr.Named("ChainReader"),
		lp:       lp,
		codec:    c,
		client:   chain.Client(),
		bindings: b,
	}, err
}

func (cr *chainReader) Name() string { return cr.lggr.Name() }

var _ commontypes.ContractTypeProvider = &chainReader{}

func (cr *chainReader) GetLatestValue(ctx context.Context, contractName, method string, params any, returnVal any) error {
	cr.lggr.Infof("!!!!!!!!!!\nEVM CR\n%s.%s\n%#v\n%s\n!!!!!!!!!!\n", contractName, method, params)
	ae, err := cr.bindings.getBinding(contractName, method, false)
	if err != nil {
		cr.lggr.Errorf("!!!!!!!!!!\nEVM CR err:\n%v\n!!!!!!!!!!\n", err)
		return err
	}

	if ae.evt == nil {
		err = cr.getLatestValueFromContract(ctx, contractName, method, params, returnVal)
		cr.lggr.Infof("!!!!!!!!!!\nEVM CR\n%s.%s\n%#v\n%s\n resp returnval %v\n!!!!!!!!!!\n", contractName, method, params, returnVal)

		return err
	}

	err = cr.getLatestValueFromLogPoller(ctx, contractName, method, *ae.evt, returnVal)

	return err
}

func (cr *chainReader) getLatestValueFromLogPoller(ctx context.Context, contractName, method string, hash common.Hash, returnVal any) error {
	cr.lggr.Infof(fmt.Sprintf("!!!!!!!!!!\nlp: EVM latest from log poller contractName: %s \n method:%s!!!!!!!!!!\n", contractName, method))
	ae, err := cr.bindings.getBinding(contractName, method, false)
	if err != nil {
		cr.lggr.Errorf("!!!!!!!!!!\nlp: EVM no binding err:\n%v\n!!!!!!!!!!\n", err)
		return err
	}

	log, err := cr.lp.LatestLogByEventSigWithConfs(hash, ae.addr, 1)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "not found") || strings.Contains(errStr, "no rows") {
			cr.lggr.Infof("!!!!!!!!!!\nlp: Returning no error when nothing is found\n!!!!!!!!!!\n")
			cr.lggr.Infof(fmt.Sprintf("!!!!!!!!!!\nlp: Returning no error when nothing is found contractName: %s \n method:%s!!!!!!!!!!\n", contractName, method))

			return nil
		}
		cr.lggr.Errorf("!!!!!!!!!!\nlp: No sig err:\n%v\n!!!!!!!!!!\n", err)
		cr.lggr.Infof(fmt.Sprintf("!!!!!!!!!!\nlp:  No sig err contractName: %s \n method:%s!!!!!!!!!!\n", contractName, method))

		return fmt.Errorf("%w: %w", commontypes.ErrInternal, err)
	}
	cr.lggr.Infof(fmt.Sprintf("!!!!!!!!!!\nlp:  No lp err  contractName: %s \n method:%s!!!!!!!!!!\n", contractName, method))

	err = cr.codec.Decode(ctx, log.Data, returnVal, wrapItemType(contractName, method, false))
	if err != nil {
		cr.lggr.Infof(fmt.Sprintf("!!!!!!!!!!\nlp: EVM decode err  contractName: %s \n method:%s!!!!!!!!!!\n", contractName, method))

		cr.lggr.Errorf("!!!!!!!!!!\nlp: EVM decode err:\n%v\n!!!!!!!!!!\n", err)
	} else {
		cr.lggr.Infof(fmt.Sprintf("!!!!!!!!!!\nlp: EVM decode success  contractName: %s \n method:%s!!!!!!!!!!\n", contractName, method))

		cr.lggr.Infof("!!!!!!!!!!\nlp: EVM decode success\n%#v\n!!!!!!!!!!\n", returnVal)
	}
	return err
}

func (cr *chainReader) getLatestValueFromContract(ctx context.Context, contractName, method string, params any, returnVal any) error {
	cr.lggr.Infof("!!!!!!!!!!\nEVM latest from contract\n!!!!!!!!!!\n")
	data, err := cr.codec.Encode(ctx, params, wrapItemType(contractName, method, true))
	if err != nil {
		cr.lggr.Errorf("!!!!!!!!!!\nEVM encode err:\n%v\n!!!!!!!!!!\n", err)
		return err
	}

	ae, err := cr.bindings.getBinding(contractName, method, true)
	if err != nil {
		cr.lggr.Errorf("!!!!!!!!!!\nEVM no binding err:\n%v\n!!!!!!!!!!\n", err)
		return err
	}
	callMsg := ethereum.CallMsg{
		To:   &ae.addr,
		From: ae.addr,
		Data: data,
	}

	output, err := cr.client.CallContract(ctx, callMsg, nil)

	if err != nil {
		cr.lggr.Errorf("!!!!!!!!!!\nEVM call err:\n%v\n!!!!!!!!!!\n", err)
		return err
	}

	cr.lggr.Infof("!!!!!!!!!!\nEVM results \n%x\n!!!!!!!!!!\n", output)

	err = cr.codec.Decode(ctx, output, returnVal, wrapItemType(contractName, method, false))
	if err != nil {
		cr.lggr.Errorf("!!!!!!!!!!\nEVM decode err:\n%v\n!!!!!!!!!!\n", err)
	} else {
		cr.lggr.Infof("!!!!!!!!!!\nEVM decode success\n%#v\n!!!!!!!!!!\n", returnVal)
	}
	return err
}

func (cr *chainReader) Start(_ context.Context) error {
	return cr.StartOnce("ChainReader", func() error {
		for contractName, contractEvents := range cr.bindings {
			for eventName, b := range contractEvents {
				if b.evt == nil {
					continue
				}

				if err := cr.lp.RegisterFilter(logpoller.Filter{
					Name:      wrapItemType(contractName, eventName, false),
					EventSigs: evmtypes.HashArray{*b.evt},
					Addresses: evmtypes.AddressArray{b.addr},
				}); err != nil {
					return fmt.Errorf("%w: %w", commontypes.ErrInternal, err)
				}
			}
		}
		return nil
	})
}
func (cr *chainReader) Close() error {
	return cr.StopOnce("ChainReader", func() error {
		for contractName, contractEvents := range cr.bindings {
			for eventName := range contractEvents {
				if err := cr.lp.UnregisterFilter(wrapItemType(contractName, eventName, false)); err != nil {
					return fmt.Errorf("%w: %w", commontypes.ErrInternal, err)
				}
			}
		}
		return nil
	})
}

func (cr *chainReader) Ready() error { return nil }
func (cr *chainReader) HealthReport() map[string]error {
	return map[string]error{cr.Name(): nil}
}

func (cr *chainReader) CreateContractType(contractName, methodName string, forEncoding bool) (any, error) {
	return cr.codec.CreateType(wrapItemType(contractName, methodName, forEncoding), forEncoding)
}

func addEventTypes(contractName, methodName string, b Bindings, contractABI abi.ABI, chainReaderDefinition types.ChainReaderDefinition, parsed *parsedTypes) error {
	event, methodExists := contractABI.Events[chainReaderDefinition.ChainSpecificName]
	if !methodExists {
		return fmt.Errorf("%w: method %s doesn't exist", commontypes.ErrInvalidConfig, chainReaderDefinition.ChainSpecificName)
	}

	if err := b.addEvent(contractName, methodName, event.ID); err != nil {
		return err
	}

	return addDecoderDef(contractName, methodName, event.Inputs, parsed, chainReaderDefinition)
}

func addMethods(
	contractName, methodName string, abi abi.ABI, chainReaderDefinition types.ChainReaderDefinition, parsed *parsedTypes) error {
	method, methodExists := abi.Methods[chainReaderDefinition.ChainSpecificName]
	if !methodExists {
		return fmt.Errorf("method: %q doesn't exist", chainReaderDefinition.ChainSpecificName)
	}

	if err := addEncoderDef(contractName, methodName, method, parsed, chainReaderDefinition); err != nil {
		return err
	}

	return addDecoderDef(contractName, methodName, method.Outputs, parsed, chainReaderDefinition)
}

func addEncoderDef(contractName, methodName string, method abi.Method, parsed *parsedTypes, chainReaderDefinition types.ChainReaderDefinition) error {
	// ABI.Pack prepends the method.ID to the encodings, we'll need the encoder to do the same.
	input := &codecEntry{Args: method.Inputs, encodingPrefix: method.ID}
	if err := input.Init(); err != nil {
		fmt.Println("encoder error isss ", err.Error())
		fmt.Printf("encoder def contract name %s \n method name %s\n method inputs %v\n err %s\n", contractName, methodName, method.Inputs, err.Error())
		return err
	}

	inputMod, err := chainReaderDefinition.InputModifications.ToModifier(evmDecoderHooks...)
	if err != nil {
		return err
	}
	input.mod = inputMod
	parsed.encoderDefs[wrapItemType(contractName, methodName, true)] = input
	return nil
}

func addDecoderDef(contractName, methodName string, outputs abi.Arguments, parsed *parsedTypes, def types.ChainReaderDefinition) error {
	output := &codecEntry{Args: outputs}
	mod, err := def.OutputModifications.ToModifier(evmDecoderHooks...)
	if err != nil {
		return err
	}
	output.mod = mod
	parsed.decoderDefs[wrapItemType(contractName, methodName, false)] = output
	if err = output.Init(); err != nil {
		fmt.Println("decoder error isss ", err.Error())
		fmt.Printf("decoder def contract name %s \n method name %s\n method outputs %v\n err: %s\n", contractName, methodName, outputs, err.Error())
	}
	return nil
}

func addTypes(chainContractReaders map[string]types.ChainContractReader, b Bindings, parsed *parsedTypes) error {
	for contractName, chainContractReader := range chainContractReaders {
		contractAbi, err := abi.JSON(strings.NewReader(chainContractReader.ContractABI))
		if err != nil {
			return err
		}

		for typeName, chainReaderDefinition := range chainContractReader.ChainReaderDefinitions {
			switch chainReaderDefinition.ReadType {
			case types.Method:
				err = addMethods(contractName, typeName, contractAbi, chainReaderDefinition, parsed)
			case types.Event:
				err = addEventTypes(contractName, typeName, b, contractAbi, chainReaderDefinition, parsed)
			default:
				return fmt.Errorf(
					"%w: invalid chain reader definition read type: %d",
					commontypes.ErrInvalidConfig,
					chainReaderDefinition.ReadType)
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func wrapItemType(contractName, methodName string, isParams bool) string {
	if isParams {
		return fmt.Sprintf("params.%s.%s", contractName, methodName)
	}
	return fmt.Sprintf("return.%s.%s", contractName, methodName)
}
