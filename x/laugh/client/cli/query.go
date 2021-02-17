package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ltacker/jupiter/x/laugh/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group laugh queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListHoho())
	cmd.AddCommand(CmdShowHoho())

	cmd.AddCommand(CmdListHihi())
	cmd.AddCommand(CmdShowHihi())

	cmd.AddCommand(CmdListHaha())
	cmd.AddCommand(CmdShowHaha())

	cmd.AddCommand(CmdListHohosent())
	cmd.AddCommand(CmdShowHohosent())

	cmd.AddCommand(CmdListHihisent())
	cmd.AddCommand(CmdShowHihisent())

	cmd.AddCommand(CmdListHahasent())
	cmd.AddCommand(CmdShowHahasent())

	return cmd
}
