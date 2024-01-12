package ocr2aggregatorchainreaderdemo

import (
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-testing-framework/utils/testcontext"

	"github.com/smartcontractkit/chainlink/v2/core/config/env"

	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/utils/ptr"

	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/utils"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/docker/test_env"
	"github.com/smartcontractkit/chainlink/integration-tests/types/config/node"
)

func validateRoundData(t *testing.T, aggregatorContracts []contracts.OffchainAggregatorV2, expectedRound int64, expectedAnswer int64) {
	for i := 0; i < len(aggregatorContracts); i++ {
		roundData, err := aggregatorContracts[i].GetLatestRound(testcontext.Get(t))
		require.NoError(t, err, "Error getting latest OCR answer")
		require.Equal(t, expectedRound, roundData.RoundId.Int64(),
			"Expected latest round from OCR2 contract to be %d but got %d", expectedRound,
			roundData.RoundId.Int64)
		require.Equal(t, expectedAnswer, roundData.Answer.Int64(),
			"Expected latest answer from OCR2 contract to be %d but got %d", expectedAnswer,
			roundData.Answer.Int64())
	}
}

func TestOCRv2BasicWithChainReaderAndCodecDemo(t *testing.T) {
	t.Parallel()
	l := logging.GetTestLogger(t)

	network, err := actions.EthereumNetworkConfigFromEnvOrDefault(l)
	require.NoError(t, err, "Error building ethereum network config")

	env, err := test_env.NewCLTestEnvBuilder().
		WithTestInstance(t).
		WithPrivateEthereumNetwork(network).
		WithMockAdapter().
		WithCLNodeConfig(node.NewConfig(node.NewBaseConfig(),
			node.WithOCR2(),
			node.WithP2Pv2(),
			node.WithTracing(),
			func(c *chainlink.Config) {
				c.Core.WebServer.HTTPMaxSize = ptr.Ptr(utils.FileSize(165536))
			},
		)).
		WithCLNodes(6).
		WithCLNodeOptions(test_env.WithNodeEnvVars(map[string]string{string(env.MedianPluginCmd): ""})).
		WithFunding(big.NewFloat(.1)).
		WithStandardCleanup().
		WithLogStream().
		Build()
	require.NoError(t, err)
	fmt.Println("Done starting")

	env.ParallelTransactions(true)

	nodeClients := env.ClCluster.NodeAPIs()
	bootstrapNode, workerNodes := nodeClients[0], nodeClients[1:]

	linkToken, err := env.ContractDeployer.DeployLinkTokenContract()
	require.NoError(t, err, "Deploying Link Token Contract shouldn't fail")

	err = actions.FundChainlinkNodesLocal(workerNodes, env.EVMClient, big.NewFloat(.05))
	require.NoError(t, err, "Error funding Chainlink nodes")

	// Gather transmitters
	var transmitters []string
	for _, node := range workerNodes {
		addr, err := node.PrimaryEthAddress()
		if err != nil {
			require.NoError(t, fmt.Errorf("error getting node's primary ETH address: %w", err))
		}
		transmitters = append(transmitters, addr)
	}

	ocrOffchainOptions := contracts.DefaultOffChainAggregatorOptions()
	aggregatorContracts, err := actions.DeployChainReaderDemoOCRv2Contracts(1, linkToken, env.ContractDeployer, transmitters, env.EVMClient, ocrOffchainOptions)
	require.NoError(t, err, "Error deploying OCRv2 aggregator contracts")

	err = actions.CreateOCRv2JobsLocal(aggregatorContracts, bootstrapNode, workerNodes, env.MockAdapter, "ocr2", 5, env.EVMClient.GetChainID().Uint64(), false, true)
	require.NoError(t, err, "Error creating OCRv2 jobs")

	ocrv2Config, err := actions.BuildMedianOCR2ConfigLocal(workerNodes, ocrOffchainOptions)
	require.NoError(t, err, "Error building OCRv2 config")

	err = actions.ConfigureOCRv2AggregatorContracts(env.EVMClient, ocrv2Config, aggregatorContracts)
	require.NoError(t, err, "Error configuring OCRv2 aggregator contracts")

	expectedAnswer := int64(5)
	for i := int64(1); i <= 100; i++ {
		require.NoError(t, actions.WatchNewOCR2Round(i, aggregatorContracts, env.EVMClient, time.Minute*10, l))
		validateRoundData(t, aggregatorContracts, i, expectedAnswer)

		i += 1
		require.NoError(t, actions.StartNewOCR2Round(i, aggregatorContracts, env.EVMClient, time.Minute*10, l))
		validateRoundData(t, aggregatorContracts, i, expectedAnswer)

		expectedAnswer += (expectedAnswer / 100) + 1
		require.NoError(t, env.MockAdapter.SetAdapterBasedIntValuePath("ocr2", []string{http.MethodGet, http.MethodPost}, int(expectedAnswer)))
	}
}
