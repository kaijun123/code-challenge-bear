package keeper

import (
	"context"
	"fmt"

	"bear/x/bear/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateBear(goCtx context.Context, msg *types.MsgUpdateBear) (*types.MsgUpdateBearResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Println("Update Bear: id desired from store: ", msg.Id)

	// TODO: Handling the message
	var bear = types.Bear{
		Role:       msg.Role,
		Background: msg.Background,
		Clothes:    msg.Clothes,
		Weapon:     msg.Weapon,
		Creator:    msg.Creator,
		Id:         msg.Id,
	}
	val, found := k.GetBear(ctx, msg.Id)
	// fmt.Println("Update Bear: id retrieved from store: ", val.Id)

	// fmt.Println("Update Bear: value of found: ", found)
	if !found {
		fmt.Println("Update Bear: failed")
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	fmt.Println("Update Bear: different creator: ", msg.Creator != val.Creator)
	if msg.Creator != val.Creator {
		fmt.Println("Update Bear: failed")
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "not creator")
	}

	k.SetBear(ctx, bear)

	// fmt.Println("Update Bear: success")
	return &types.MsgUpdateBearResponse{}, nil
}
