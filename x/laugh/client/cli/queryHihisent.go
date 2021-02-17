package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ltacker/jupiter/x/laugh/types"
	"github.com/spf13/cobra"
)

func CmdListHihisent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-hihisent",
		Short: "list all hihisent",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHihisentRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HihisentAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowHihisent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-hihisent [id]",
		Short: "shows a hihisent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetHihisentRequest{
				Id: args[0],
			}

			res, err := queryClient.Hihisent(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
