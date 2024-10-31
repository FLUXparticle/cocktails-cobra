package cocktails

import (
	"cocktails-cobra/pkg/cocktails/server"
	"github.com/spf13/cobra"
)

var useChi bool
var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"serv"},
	Short:   "Runs server",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer(useChi)
	},
}

func init() {
	serverCmd.Flags().BoolVarP(&useChi, "chi", "c", false, "Use Chi")
	rootCmd.AddCommand(serverCmd)
}
