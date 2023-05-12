package rand_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "rand/testutil/keeper"
	"rand/testutil/nullify"
	"rand/x/rand"
	"rand/x/rand/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RandKeeper(t)
	rand.InitGenesis(ctx, *k, genesisState)
	got := rand.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
