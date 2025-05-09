package cocktails

import (
	"cocktails-cobra/pkg/cocktails/repository"
	"cocktails-cobra/pkg/cocktails/repository/database"
	"cocktails-cobra/pkg/cocktails/repository/local"
	"cocktails-cobra/pkg/cocktails/repository/remote"
	"cocktails-cobra/pkg/cocktails/service"
	"fmt"
)

// initService initializes the cocktail service with the appropriate repository
// based on the flags and returns it
func initService(offline bool, useDatabase bool) *service.CocktailService {
	var repo repository.CocktailRepository

	if useDatabase {
		// Initialize the database and create a database repository
		db := database.NewDatabase()
		repo = database.NewDatabaseRepository(db)
		fmt.Println("Running with database repository")
	} else if offline {
		repo = local.NewLocalRepository()
		fmt.Println("Running in offline mode with local repository")
	} else {
		repo = remote.NewRemoteRepository()
		fmt.Println("Running in online mode with remote repository")
	}

	return service.NewCocktailService(repo)
}
