package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	epitomeclcmd "github.com/jsyun-epi/cosmos-app-test/x/epitomecl/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group innsa queries under a subcommand
	namesvcQueryCmd := &cobra.Command{
		Use:   "epitomecl",
		Short: "Querying commands for the epitomecl module",
	}

	namesvcQueryCmd.AddCommand(client.GetCommands(
		epitomeclcmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		epitomeclcmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return namesvcQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	namesvcTxCmd := &cobra.Command{
		Use:   "epitomecl",
		Short: "epitomecl transactions subcommands",
	}

	namesvcTxCmd.AddCommand(client.PostCommands(
		epitomeclcmd.GetCmdBuyName(mc.cdc),
		epitomeclcmd.GetCmdSetName(mc.cdc),
	)...)

	return namesvcTxCmd
}


