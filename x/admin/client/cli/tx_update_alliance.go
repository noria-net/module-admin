package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/noria-net/module-admin/x/admin/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateAlliance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-alliance [denom] [reward-weight] [consensus-weight] [consensus-cap]",
		Short: "Broadcast message update-alliance",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argRewardWeight := args[1]
			argConsensusWeight := args[2]
			argConsensusCap := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAlliance(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argRewardWeight,
				argConsensusWeight,
				argConsensusCap,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
