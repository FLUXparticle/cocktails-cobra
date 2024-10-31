package cocktails

import (
	"cocktails-cobra/pkg/cocktails/initdb"
	"github.com/spf13/cobra"
)

var initDbCmd = &cobra.Command{
	Use:     "initdb",
	Aliases: []string{"init"},
	Short:   "Initialize the database",
	Run: func(cmd *cobra.Command, args []string) {
		initdb.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(initDbCmd)
}
