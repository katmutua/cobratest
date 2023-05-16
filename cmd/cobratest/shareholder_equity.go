package cobratest

import (
	"fmt"

	"github.com/katmutua/cobratest/pkg/cobratest"
	"github.com/spf13/cobra"
)

var equityCmd = &cobra.Command{
	Use:     "equity",
	Aliases: []string{"eq"},
	Short:   "Shows how to calculate the shareholder equity",
	Long:    "Shows a complete formula on how to calculate the shareholder equity",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		result := cobratest.ShareHolderEquity()
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(equityCmd)
}
