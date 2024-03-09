package keeper

import (
	"context"

	"bear/x/bear/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBear(goCtx context.Context, msg *types.MsgCreateBear) (*types.MsgCreateBearResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateBearResponse{}, nil
}
