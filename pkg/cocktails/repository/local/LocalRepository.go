package local

import (
	"cocktails-cobra/pkg/cocktails/model"
	"errors"
)

// LocalRepository implements CocktailRepository by providing data from a local source
type LocalRepository struct {
	// This could be a local database, file storage, or in-memory cache
	cocktails map[uint]*model.Cocktail
}

// NewLocalRepository creates a new instance of LocalRepository
func NewLocalRepository() *LocalRepository {
	// Initialize with some sample data
	// In a real implementation, this might load from a local file or database
	return &LocalRepository{
		cocktails: make(map[uint]*model.Cocktail),
	}
}

// GetCocktails returns all cocktails from the local storage
func (r *LocalRepository) GetCocktails() ([]*model.Cocktail, error) {
	if len(r.cocktails) == 0 {
		return nil, errors.New("no cocktails available in local storage")
	}

	// Convert map to slice
	cocktails := make([]*model.Cocktail, 0, len(r.cocktails))
	for _, cocktail := range r.cocktails {
		cocktails = append(cocktails, cocktail)
	}

	return cocktails, nil
}

// GetCocktail returns a specific cocktail by ID from the local storage
func (r *LocalRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	cocktail, exists := r.cocktails[id]
	if !exists {
		return nil, errors.New("cocktail not found in local storage")
	}

	return cocktail, nil
}

// AddCocktail adds a cocktail to the local storage
// This is useful for caching cocktails fetched from the remote repository
func (r *LocalRepository) AddCocktail(cocktail *model.Cocktail) {
	r.cocktails[cocktail.CocktailID] = cocktail
}
