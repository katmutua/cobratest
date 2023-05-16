package cobratest

import (
	"fmt"

	"github.com/katmutua/cobratest/pkg/cobratest"
	"github.com/spf13/cobra"
)

var assetsCmd = &cobra.Command{
	Use:     "assets",
	Aliases: []string{"as"},
	Short:   "",
	Long:    "",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		result := cobratest.Assets()
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(assetsCmd)
}
