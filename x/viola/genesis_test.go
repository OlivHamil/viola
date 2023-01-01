package viola_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "viola/testutil/keeper"
	"viola/testutil/nullify"
	"viola/x/viola"
	"viola/x/viola/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ViolaKeeper(t)
	viola.InitGenesis(ctx, *k, genesisState)
	got := viola.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
