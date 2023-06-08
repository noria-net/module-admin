package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorslegacy "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAlliance = "update_alliance"

var _ sdk.Msg = &MsgUpdateAlliance{}

func NewMsgUpdateAlliance(creator string, denom string, rewardWeight string, consensusWeight string, consensusCap string) *MsgUpdateAlliance {
	return &MsgUpdateAlliance{
		Creator:         creator,
		Denom:           denom,
		RewardWeight:    rewardWeight,
		ConsensusWeight: consensusWeight,
		ConsensusCap:    consensusCap,
	}
}

func (msg *MsgUpdateAlliance) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAlliance) Type() string {
	return TypeMsgUpdateAlliance
}

func (msg *MsgUpdateAlliance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAlliance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAlliance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorslegacy.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// assert if msg.RewardWeight can be parsed to sdk.Dec
	_, err = sdk.NewDecFromStr(msg.RewardWeight)
	if err != nil {
		return sdkerrors.Wrapf(ErrUpdateAlliance, "invalid reward weight: %s", msg.RewardWeight)
	}

	// assert if msg.ConsensusWeight can be parsed to sdk.Dec
	_, err = sdk.NewDecFromStr(msg.ConsensusWeight)
	if err != nil {
		return sdkerrors.Wrapf(ErrUpdateAlliance, "invalid consensus weight: %s", msg.ConsensusWeight)
	}

	// assert if msg.ConsensusCap can be parsed to sdk.Dec
	_, err = sdk.NewDecFromStr(msg.ConsensusCap)
	if err != nil {
		return sdkerrors.Wrapf(ErrUpdateAlliance, "invalid consensus cap: %s", msg.ConsensusCap)
	}
	return nil
}
