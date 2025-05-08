package repository

import (
	"cocktails-cobra/pkg/cocktails/model"
)

// MutableCocktailRepository extends CocktailRepository with editing functions
type MutableCocktailRepository interface {
	// Embed the CocktailRepository interface to inherit its methods
	CocktailRepository

	// AddCocktail adds a cocktail to the repository
	AddCocktail(cocktail *model.Cocktail)
}
