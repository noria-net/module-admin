package keeper_test

import (
	"testing"

	testkeeper "github.com/noria-net/module-admin/testutil/keeper"
	"github.com/noria-net/module-admin/x/admin/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AdminKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Admin, k.Admin(ctx))
}
