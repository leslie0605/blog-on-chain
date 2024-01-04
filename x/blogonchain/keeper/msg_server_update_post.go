package keeper

import (
	"context"

	"blog-on-chain/x/blogonchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdatePostResponse{}, nil
}
