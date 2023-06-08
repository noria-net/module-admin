package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	alliancetypes "github.com/noria-net/alliance/x/alliance/types"
)

type AllianceKeeper interface {
	UpdateAllianceAsset(ctx sdk.Context, newAsset alliancetypes.AllianceAsset) error
	GetAssetByDenom(ctx sdk.Context, denom string) (asset alliancetypes.AllianceAsset, found bool)
	// Methods imported from alliance should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
