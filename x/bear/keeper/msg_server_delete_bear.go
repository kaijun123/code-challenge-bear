package keeper

import (
	"context"
	"fmt"

	"bear/x/bear/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteBear(goCtx context.Context, msg *types.MsgDeleteBear) (*types.MsgDeleteBearResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	id := msg.Id
	val, found := k.GetBear(ctx, id)

	// fmt.Println("Delete Bear: id desired from store: ", msg.Id)
	// fmt.Println("Delete Bear: id retrieved from store: ", val.Id)

	// fmt.Println("Delete Bear: creator desired from store: ", msg.Creator)
	// fmt.Println("Delete Bear: creator retrieved from store:  ", val.Creator)

	// fmt.Println("Delete Bear: value of found: ", found)
	if !found {
		// fmt.Println("Delete Bear: failed")

		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	// fmt.Println("Delete Bear: different creator: ", msg.Creator != val.Creator)
	if msg.Creator != val.Creator {
		// fmt.Println("Delete Bear: failed")
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "not creator")
	}

	k.RemoveBear(ctx, id)

	// fmt.Println("Delete Bear: success")
	return &types.MsgDeleteBearResponse{}, nil
}
