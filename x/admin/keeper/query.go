package keeper

import (
	"github.com/noria-net/module-admin/x/admin/types"
)

var _ types.QueryServer = Keeper{}
