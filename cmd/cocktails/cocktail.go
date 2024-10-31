package cocktails

import (
	"cocktails-cobra/pkg/cocktails"
	"github.com/spf13/cobra"
)

var cocktailID uint

var cocktailCmd = &cobra.Command{
	Use:   "cocktail",
	Short: "Zeigt alle Cocktails an oder einen bestimmten Cocktail, wenn die ID via Flag gesetzt wird.",
	Long:  "Nutze das Flag --id (oder -i), um einen bestimmten Cocktail anzuzeigen. Wird die ID nicht gesetzt, listet der Befehl alle Cocktails auf.",
	Run: func(cmd *cobra.Command, args []string) {
		if cocktailID > 0 {
			// Ein spezieller Cocktail wird angefragt.
			// Hier übergeben wir die ID als int, in dem wir sie, falls nötig, in einen String konvertieren.
			cocktails.DisplayCocktail(cocktailID)
		} else {
			// Wenn keine ID gesetzt ist, werden alle Cocktails abgefragt.
			cocktails.DisplayCocktails()
		}
	},
}

func init() {
	rootCmd.AddCommand(cocktailCmd)
	// Flag für die Cocktail ID, optional. Standard ist 0.
	cocktailCmd.Flags().UintVarP(&cocktailID, "id", "i", 0, "ID des Cocktails (optional)")
}
