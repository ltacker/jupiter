package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ltacker/jupiter/x/laugh/types"
	"github.com/spf13/cobra"
)

func CmdListHaha() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-haha",
		Short: "list all haha",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHahaRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HahaAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowHaha() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-haha [id]",
		Short: "shows a haha",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetHahaRequest{
				Id: args[0],
			}

			res, err := queryClient.Haha(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
