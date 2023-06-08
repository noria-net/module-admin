package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorslegacy "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noria-net/module-admin/x/admin/types"
)

func (k msgServer) UpdateAlliance(goCtx context.Context, msg *types.MsgUpdateAlliance) (*types.MsgUpdateAllianceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	err = k.ValidateAdmin(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrorslegacy.ErrUnauthorized, "invalid admin: %s", err.Error())
	}

	asset, found := k.allianceKeeper.GetAssetByDenom(ctx, msg.Denom)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUpdateAlliance, "Asset with denom: %s does not exist", msg.Denom)
	}

	asset.RewardWeight = sdk.MustNewDecFromStr(msg.RewardWeight)
	asset.ConsensusWeight = sdk.MustNewDecFromStr(msg.ConsensusWeight)
	asset.ConsensusCap = sdk.MustNewDecFromStr(msg.ConsensusCap)

	err = k.allianceKeeper.UpdateAllianceAsset(ctx, asset)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUpdateAlliance, "could not update alliance asset: %s", err.Error())
	}

	return &types.MsgUpdateAllianceResponse{}, nil
}
