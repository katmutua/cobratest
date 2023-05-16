package cobratest

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/katmutua/cobratest/pkg/cobratest"
)

var equityCmd = &cobra.Command{
	Use:   "equity",
	Aliases: []string{"eq"},
	Short: "Shows how to calculate the shareholder equity",
	Long:  "Shows a complete formula on how to calculate the shareholder equity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := cobratest.shareHolderEquity(args[0])
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(equityCmd)
}
