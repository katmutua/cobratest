package cobratest

import (
	"fmt"

	"github.com/katmutua/cobratest/pkg/cobratest"
	"github.com/spf13/cobra"
)

var totalLiabilities = &cobra.Command{
	Use:     "totalLiabilities",
	Aliases: []string{"tl"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// value := args[0]
		result := cobratest.TotalLiabilities()
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(totalLiabilities)
}
