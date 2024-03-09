package keeper

import (
	"context"

	"bear/x/bear/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateBear(goCtx context.Context, msg *types.MsgUpdateBear) (*types.MsgUpdateBearResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateBearResponse{}, nil
}
