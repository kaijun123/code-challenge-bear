package bear_test

import (
	"testing"

	keepertest "bear/testutil/keeper"
	"bear/testutil/nullify"
	bear "bear/x/bear/module"
	"bear/x/bear/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BearKeeper(t)
	bear.InitGenesis(ctx, k, genesisState)
	got := bear.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
