package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleRapName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgSetRapName struct {
	RapName string         `json:"rapname"`
	Value   string         `json:"value"`
	Owner   sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetRapName(rapname string, value string, owner sdk.AccAddress) MsgSetRapName {
	return MsgSetName{
		Name:  rapname,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetRapName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetRapName) Type() string { return "set_rapname" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetRapName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.RapName) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("RapName and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetRapName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetRapName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyName defines the BuyName message
type MsgBuyRapName struct {
	Name  string         `json:"rapname"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyRapName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyRapName {
	return MsgBuyName{
		RapName: rapname,
		Bid:     bid,
		Buyer:   buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyRapName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyRapName) Type() string { return "buy_rapname" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyRApName) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.RapName) == 0 {
		return sdk.ErrUnknownRequest("RapName cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyRapName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyRapName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// MsgDeleteName defines a DeleteName message
type MsgDeleteRapName struct {
	RapName string         `json:"rapname"`
	Owner   sdk.AccAddress `json:"owner"`
}

// NewMsgDeleteName is a constructor function for MsgDeleteName
func NewMsgDeleteRapName(rapname string, owner sdk.AccAddress) MsgDeleteRapName {
	return MsgDeleteRapName{
		RapName: rapname,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteRapName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteRapName) Type() string { return "delete_rapname" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteRapName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.RapName) == 0 {
		return sdk.ErrUnknownRequest("RapName cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteRapName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
