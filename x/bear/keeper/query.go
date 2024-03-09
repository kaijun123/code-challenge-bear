package keeper

import (
	"bear/x/bear/types"
)

var _ types.QueryServer = Keeper{}
