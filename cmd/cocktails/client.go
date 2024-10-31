package cocktails

import (
	"cocktails-cobra/pkg/cocktails/client"
	"fmt"
	"github.com/spf13/cobra"
)

var runParallel bool
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Runs client",
	Run: func(cmd *cobra.Command, args []string) {
		var sumMilk int
		if runParallel {
			sumMilk = client.DoParallel()
		} else {
			sumMilk = client.DoSequential()
		}
		fmt.Printf("sumMilk: %d\n", sumMilk)
	},
}

func init() {
	clientCmd.Flags().BoolVarP(&runParallel, "parallel", "p", false, "Run parallel")
	rootCmd.AddCommand(clientCmd)
}
