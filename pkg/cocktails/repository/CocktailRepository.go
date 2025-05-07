package repository

import (
	"cocktails-cobra/pkg/cocktails/model"
)

// CocktailRepository defines the interface for cocktail data access
type CocktailRepository interface {
	// GetCocktails returns all available cocktails
	GetCocktails() ([]*model.Cocktail, error)

	// GetCocktail returns a specific cocktail by ID
	GetCocktail(id uint) (*model.Cocktail, error)
}
