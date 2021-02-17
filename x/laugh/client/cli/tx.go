package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ltacker/jupiter/x/laugh/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdSendHohomes())

	cmd.AddCommand(CmdSendHihimes())

	cmd.AddCommand(CmdSendHahames())

	cmd.AddCommand(CmdCreateHoho())
	cmd.AddCommand(CmdUpdateHoho())
	cmd.AddCommand(CmdDeleteHoho())

	cmd.AddCommand(CmdCreateHihi())
	cmd.AddCommand(CmdUpdateHihi())
	cmd.AddCommand(CmdDeleteHihi())

	cmd.AddCommand(CmdCreateHaha())
	cmd.AddCommand(CmdUpdateHaha())
	cmd.AddCommand(CmdDeleteHaha())

	cmd.AddCommand(CmdCreateHohosent())
	cmd.AddCommand(CmdUpdateHohosent())
	cmd.AddCommand(CmdDeleteHohosent())

	cmd.AddCommand(CmdCreateHihisent())
	cmd.AddCommand(CmdUpdateHihisent())
	cmd.AddCommand(CmdDeleteHihisent())

	cmd.AddCommand(CmdCreateHahasent())
	cmd.AddCommand(CmdUpdateHahasent())
	cmd.AddCommand(CmdDeleteHahasent())

	return cmd
}
