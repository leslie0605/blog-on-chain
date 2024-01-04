package keeper

import (
	"blog-on-chain/x/blogonchain/types"
)

var _ types.QueryServer = Keeper{}
