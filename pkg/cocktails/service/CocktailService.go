package service

import (
	"cocktails-cobra/pkg/cocktails/repository"
	"fmt"
)

// CocktailService provides business logic for cocktail operations
type CocktailService struct {
	repo repository.CocktailRepository
}

// NewCocktailService creates a new instance of CocktailService with the provided repository
func NewCocktailService(repo repository.CocktailRepository) *CocktailService {
	return &CocktailService{
		repo: repo,
	}
}

// DisplayCocktails prints all cocktails to the console
// This is a convenience method for CLI applications
func (s *CocktailService) DisplayCocktails() error {
	cocktails, err := s.repo.GetCocktails()
	if err != nil {
		return fmt.Errorf("error getting cocktails: %w", err)
	}

	fmt.Println("Available Cocktails:")
	for _, cocktail := range cocktails {
		fmt.Printf("- %d: %s\n", cocktail.CocktailID, cocktail.Name)
	}

	return nil
}

// DisplayCocktail prints a specific cocktail to the console
// This is a convenience method for CLI applications
func (s *CocktailService) DisplayCocktail(id uint) error {
	cocktail, err := s.repo.GetCocktail(id)
	if err != nil {
		return fmt.Errorf("error getting cocktail with ID %d: %w", id, err)
	}

	fmt.Printf("Cocktail: %s (ID: %d)\n", cocktail.Name, cocktail.CocktailID)
	fmt.Println("Instructions:")
	for i, instruction := range cocktail.Instructions {
		fmt.Printf("%d. %v\n", i+1, instruction)
	}

	return nil
}
