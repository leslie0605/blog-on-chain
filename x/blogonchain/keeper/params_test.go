package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "blog-on-chain/testutil/keeper"
	"blog-on-chain/x/blogonchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BlogonchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
