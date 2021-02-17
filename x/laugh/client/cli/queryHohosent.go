package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ltacker/jupiter/x/laugh/types"
	"github.com/spf13/cobra"
)

func CmdListHohosent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-hohosent",
		Short: "list all hohosent",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHohosentRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HohosentAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowHohosent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-hohosent [id]",
		Short: "shows a hohosent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetHohosentRequest{
				Id: args[0],
			}

			res, err := queryClient.Hohosent(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
