package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/values"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

const testID = "test-id-1"

func TestOnDemandTrigger(t *testing.T) {
	r := capabilities.NewRegistry()
	tr := NewOnDemandTrigger()
	ctx := testutils.Context(t)

	err := r.Add(ctx, tr)
	require.NoError(t, err)

	trigger, err := r.GetTrigger(ctx, tr.Info().ID)
	require.NoError(t, err)

	callback := make(chan capabilities.CapabilityResponse, 10)

	req := capabilities.CapabilityRequest{
		Metadata: capabilities.RequestMetadata{
			WorkflowExecutionID: testID,
		},
	}
	err = trigger.RegisterTrigger(ctx, callback, req)
	require.NoError(t, err)

	er := capabilities.CapabilityResponse{
		Value: &values.String{Underlying: testID},
	}

	err = tr.FanOutEvent(ctx, er)
	require.NoError(t, err)

	assert.Len(t, callback, 1)
	assert.Equal(t, er, <-callback)
}

func TestOnDemandTrigger_ChannelDoesntExist(t *testing.T) {
	tr := NewOnDemandTrigger()
	ctx := testutils.Context(t)

	er := capabilities.CapabilityResponse{
		Value: &values.String{Underlying: testID},
	}
	err := tr.SendEvent(ctx, testID, er)
	assert.ErrorContains(t, err, "no registration")
}
