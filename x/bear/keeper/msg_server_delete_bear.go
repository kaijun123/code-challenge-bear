package keeper

import (
	"context"

	"bear/x/bear/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteBear(goCtx context.Context, msg *types.MsgDeleteBear) (*types.MsgDeleteBearResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteBearResponse{}, nil
}
