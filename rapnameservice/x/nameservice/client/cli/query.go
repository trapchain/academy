package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/sdk-tutorials/raprapnameservice/x/raprapnameservice/internal/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	raprapnameserviceQueryCmd := &cobra.Command{
		Use:                        types.Modulerapname,
		Short:                      "Querying commands for the raprapnameservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	raprapnameserviceQueryCmd.AddCommand(client.GetCommands(
		GetCmdResolverapname(storeKey, cdc),
		GetCmdWhois(storeKey, cdc),
		GetCmdrapnames(storeKey, cdc),
	)...)
	return raprapnameserviceQueryCmd
}

// GetCmdResolverapname queries information about a rapname
func GetCmdResolverapname(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve [rapname]",
		Short: "resolve rapname",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			rapname := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, rapname), nil)
			if err != nil {
				fmt.Printf("could not resolve rapname - %s \n", rapname)
				return nil
			}

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdWhois queries information about a domain
func GetCmdWhois(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "whois [rapname]",
		Short: "Query whois info of rapname",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			rapname := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, rapname), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", rapname)
				return nil
			}

			var out types.Whois
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdrapnames queries a list of all rapnames
func GetCmdrapnames(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "rapnames",
		Short: "rapnames",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/rapnames", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query rapnames\n")
				return nil
			}

			var out types.QueryResrapnames
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
