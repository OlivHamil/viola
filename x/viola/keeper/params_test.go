package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "viola/testutil/keeper"
	"viola/x/viola/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ViolaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
