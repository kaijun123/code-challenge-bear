package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBear{}

func NewMsgCreateBear(creator string, role string, background string, clothes string, weapon string) *MsgCreateBear {
	return &MsgCreateBear{
		Creator:    creator,
		Role:       role,
		Background: background,
		Clothes:    clothes,
		Weapon:     weapon,
	}
}

func (msg *MsgCreateBear) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
