package keeper

import (
	"context"

	"bear/x/bear/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBear(goCtx context.Context, msg *types.MsgCreateBear) (*types.MsgCreateBearResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var bear = types.Bear{
		Role:       msg.Role,
		Background: msg.Background,
		Clothes:    msg.Clothes,
		Weapon:     msg.Weapon,
		Creator:    msg.Creator,
	}

	id := k.AppendBear(
		ctx,
		bear,
	)
	return &types.MsgCreateBearResponse{
		Id: id,
	}, nil
}
