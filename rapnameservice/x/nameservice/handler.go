package nameservice

import (
	"fmt"

	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetRapName:
			return handleMsgSetRapName(ctx, keeper, msg)
		case MsgBuyRapName:
			return handleMsgBuyRapName(ctx, keeper, msg)
		case MsgDeleteRapName:
			return handleMsgDeleteRapName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized rapnameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetRapName(ctx sdk.Context, keeper Keeper, msg MsgSetRapName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.RapName)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetRapName(ctx, msg.RapName, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                            // return
}

// Handle a message to buy name
func handleMsgBuyRapName(ctx sdk.Context, keeper Keeper, msg MsgBuyRapName) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.RapName).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.RapName) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.RapName, msg.Buyer)
	keeper.SetPrice(ctx, msg.RapName, msg.Bid)
	return sdk.Result{}
}

// Handle a message to delete name
func handleMsgDeleteRapName(ctx sdk.Context, keeper Keeper, msg MsgDeleteRapName) sdk.Result {
	if !keeper.IsRapNamePresent(ctx, msg.Name) {
		return types.ErrRapNameDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.RapName)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteWhois(ctx, msg.RapName)
	return sdk.Result{}
}
