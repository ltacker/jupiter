package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ltacker/jupiter/x/laugh/types"
	"github.com/spf13/cobra"
)

func CmdListHihi() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-hihi",
		Short: "list all hihi",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHihiRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HihiAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowHihi() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-hihi [id]",
		Short: "shows a hihi",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetHihiRequest{
				Id: args[0],
			}

			res, err := queryClient.Hihi(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
