package remote

import (
	"cocktails-cobra/pkg/cocktails/model"
	"errors"
)

const baseURL string = "https://cocktails.fluxparticle.com"

// RemoteRepository implements CocktailRepository by fetching data from a remote API
type RemoteRepository struct {
}

// NewRemoteRepository creates a new instance of RemoteRepository
func NewRemoteRepository() *RemoteRepository {
	return &RemoteRepository{}
}

// GetCocktails fetches all cocktails from the remote API
func (r *RemoteRepository) GetCocktails() ([]*model.Cocktail, error) {
	// TODO: Implement actual API call to fetch cocktails
	// This would typically involve making an HTTP request to the API
	// and parsing the response into model.Cocktail objects

	// For now, return a placeholder implementation
	return nil, errors.New("remote repository GetCocktails not implemented")
}

// GetCocktail fetches a specific cocktail by ID from the remote API
func (r *RemoteRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	// TODO: Implement actual API call to fetch a specific cocktail
	// This would typically involve making an HTTP request to the API with the ID
	// and parsing the response into a model.Cocktail object

	// For now, return a placeholder implementation
	return nil, errors.New("remote repository GetCocktail not implemented")
}
