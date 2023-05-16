package cobratest

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   `cobratest`,
	Short: `cobratest - a simple cli to show accounting formulas`,
	Long: `cobratest is a super fancy cli for money management ... trust it over your accountant (just kidding)
One can use cobratest to list basic accounting formulae`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. The sounmd of a million screaming accountants due to failure to run cmd '&s'", err)
		os.Exit(1)
	}
}
