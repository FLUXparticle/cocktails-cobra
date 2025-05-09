package cocktails

import (
	"cocktails-cobra/pkg/cocktails/handler"
	"cocktails-cobra/pkg/cocktails/repository"
	"cocktails-cobra/pkg/cocktails/repository/database"
	"cocktails-cobra/pkg/cocktails/repository/local"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Startet einen HTTP-Server für Cocktail-Informationen",
	Long:  "Startet einen HTTP-Server, der eine REST-API für Cocktail-Informationen bereitstellt. Der Server läuft, bis er manuell beendet wird.",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new Gin router
		router := gin.Default()

		// Initialize the database and create a database repository
		var repo repository.MutableCocktailRepository

		if useDatabase {
			db := database.NewDatabase()
			repo = database.NewDatabaseRepository(db)
			fmt.Println("Running with database repository")
		} else {
			repo = local.NewLocalRepository()
			fmt.Println("Running with memory repository")
		}

		// Create a new CocktailHandler with the repository from the service
		cocktailHandler := handler.NewCocktailHandler(repo)

		// Register the routes
		cocktailHandler.RegisterRoutes(router)

		// Print server information
		fmt.Println("Starting server on http://localhost:8080")
		fmt.Println("Press Ctrl+C to stop the server")

		// Start the server
		router.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
