package blogonchain_test

import (
	"testing"

	keepertest "blog-on-chain/testutil/keeper"
	"blog-on-chain/testutil/nullify"
	"blog-on-chain/x/blogonchain/module"
	"blog-on-chain/x/blogonchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogonchainKeeper(t)
	blogonchain.InitGenesis(ctx, k, genesisState)
	got := blogonchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
