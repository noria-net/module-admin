package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorslegacy "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/noria-net/module-admin/x/admin/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		allianceKeeper types.AllianceKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	allianceKeeper types.AllianceKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		allianceKeeper: allianceKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ValidateAdmin(ctx sdk.Context, admin string) error {
	moduleAdmin := k.Admin(ctx)
	if len(moduleAdmin) == 0 {
		return types.ErrAdminNotInitialized
	}
	_, err := sdk.AccAddressFromBech32(admin)
	if err != nil {
		return sdkerrorslegacy.ErrInvalidAddress
	}
	if k.Admin(ctx) != admin {
		return sdkerrorslegacy.ErrUnauthorized
	}
	return nil
}
