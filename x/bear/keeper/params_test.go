package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bear/testutil/keeper"
	"bear/x/bear/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BearKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
