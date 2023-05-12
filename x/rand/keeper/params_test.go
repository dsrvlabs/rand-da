package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "rand/testutil/keeper"
	"rand/x/rand/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RandKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
