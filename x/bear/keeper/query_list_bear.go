package keeper

import (
	"context"

	"bear/x/bear/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListBear(goCtx context.Context, req *types.QueryListBearRequest) (*types.QueryListBearResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BearKey))

	var bears []types.Bear
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var bear types.Bear
		if err := k.cdc.Unmarshal(value, &bear); err != nil {
			return err
		}

		bears = append(bears, bear)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListBearResponse{Bear: bears, Pagination: pageRes}, nil
}
