package cocktails

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cocktails",
	Short: "cocktails - a simple CLI app for cocktails",
	Long: `Cocktails CLI - Eine Kommandozeilenanwendung zur Verwaltung und Anzeige von Cocktail-Informationen.

Diese Anwendung ermöglicht es Ihnen, Informationen über Cocktails und deren Zutaten abzurufen sowie den Status Ihres virtuellen Kühlschranks zu verwalten. 

Verwenden Sie die verschiedenen Unterbefehle, um auf die spezifischen Funktionen zuzugreifen. Jeder Unterbefehl bietet seine eigene Hilfe und Optionen, die Sie mit dem Flag --help anzeigen können.

Die Anwendung kommuniziert mit einer externen API, um die Daten abzurufen und zu aktualisieren.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
