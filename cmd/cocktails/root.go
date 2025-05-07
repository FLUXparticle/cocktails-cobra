package cocktails

import (
	"cocktails-cobra/pkg/cocktails/service"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Global flags and service
var (
	offlineMode bool
	// cocktailService is the global service instance used by all commands
	cocktailService *service.CocktailService
)

var rootCmd = &cobra.Command{
	Use:   "cocktails",
	Short: "cocktails - a simple CLI app for cocktails",
	Long: `Cocktails CLI - Eine Kommandozeilenanwendung zur Verwaltung und Anzeige von Cocktail-Informationen.

Diese Anwendung ermöglicht es Ihnen, Informationen über Cocktails und deren Zutaten abzurufen sowie den Status Ihres virtuellen Kühlschranks zu verwalten. 

Verwenden Sie die verschiedenen Unterbefehle, um auf die spezifischen Funktionen zuzugreifen. Jeder Unterbefehl bietet seine eigene Hilfe und Optionen, die Sie mit dem Flag --help anzeigen können.

Die Anwendung kommuniziert mit einer externen API, um die Daten abzurufen und zu aktualisieren.`,
	// Initialize the cocktail service before any command runs
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize the cocktail service with the appropriate repository
		// based on the offline flag and store it in the global variable
		cocktailService = initService(offlineMode)
	},
}

func init() {
	// Define global flags
	rootCmd.PersistentFlags().BoolVar(&offlineMode, "offline", false, "Run in offline mode using local data")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
