package rapnameservice

import (
	"github.com/cosmos/sdk-tutorials/rapnameservice/x/rapnameservice/internal/keeper"
	"github.com/cosmos/sdk-tutorials/rapnameservice/x/rapnameservice/internal/types"
)

const (
	ModuleRapName = types.ModuleRapName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyRapName    = types.NewMsgBuyRapName
	NewMsgSetRapName    = types.NewMsgSetapName
	NewMsgDeleteRapName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	MsgSetRapName      = types.MsgSetRapName
	MsgBuyRapName      = types.MsgBuyRapName
	MsgDeleteRapName   = types.MsgDeleteRapName
	QueryResResolve = types.QueryResResolve
	QueryResRapNames   = types.QueryResRapNames
	Whois           = types.Whois
)
