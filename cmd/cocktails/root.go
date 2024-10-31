package cocktails

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cocktails",
	Short: "cocktails - a simple CLI server for cocktails",
	Long: `cocktail API
Server: http://192.168.82.204:8080

Endpoint: GET /cocktails
{
    "0": "Cocktail 1",
    "1": "Cocktail 2",
    "2": "Cocktail 3",
    ...
}

Endpoint: GET /cocktails/{id}
{
    "name": "Cocktail Name",
    "instructions": [
        { "cl": 50, "ingredient": "ABC" },
        { "cl": 20, "ingredient": "EFG" },
    ]
}`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
