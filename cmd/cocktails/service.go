package cocktails

import (
	"cocktails-cobra/pkg/cocktails/repository"
	"cocktails-cobra/pkg/cocktails/repository/local"
	"cocktails-cobra/pkg/cocktails/repository/remote"
	"cocktails-cobra/pkg/cocktails/service"
	"fmt"
)

// initService initializes the cocktail service with the appropriate repository
// based on the offline flag and returns it
func initService(offline bool) *service.CocktailService {
	var repo repository.CocktailRepository

	if offline {
		repo = local.NewLocalRepository()
		fmt.Println("Running in offline mode with local repository")
	} else {
		repo = remote.NewRemoteRepository()
		fmt.Println("Running in online mode with remote repository")
	}

	return service.NewCocktailService(repo)
}
