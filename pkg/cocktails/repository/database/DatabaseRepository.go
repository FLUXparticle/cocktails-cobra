package database

import (
	"cocktails-cobra/pkg/cocktails/model"
	"cocktails-cobra/pkg/cocktails/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DatabaseRepository implements MutableCocktailRepository using GORM for database operations
type DatabaseRepository struct {
	db *gorm.DB
}

// NewDatabaseRepository creates a new instance of DatabaseRepository
// Ensure DatabaseRepository implements repository.MutableCocktailRepository
func NewDatabaseRepository(db *gorm.DB) repository.MutableCocktailRepository {
	// Auto migrate the schema
	db.AutoMigrate(&model.Cocktail{}, &model.Instruction{}, &model.Ingredient{})

	return &DatabaseRepository{
		db: db,
	}
}

// NewDatabase initializes and returns a new GORM database connection
func NewDatabase() *gorm.DB {
	cfg := &gorm.Config{}
	dialector := sqlite.Open("cocktails.db")
	db, err := gorm.Open(dialector, cfg)
	if err != nil {
		panic(err)
	}
	return db
}

// GetCocktails returns all cocktails from the database
func (r *DatabaseRepository) GetCocktails() ([]*model.Cocktail, error) {
	var cocktails []*model.Cocktail

	result := r.db.Find(&cocktails)
	if result.Error != nil {
		return nil, result.Error
	}

	return cocktails, nil
}

// GetCocktail returns a specific cocktail by ID from the database
func (r *DatabaseRepository) GetCocktail(id uint) (*model.Cocktail, error) {
	var cocktail model.Cocktail

	result := r.db.Preload("Instructions.Ingredient").First(&cocktail, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cocktail, nil
}

// AddCocktail adds a cocktail to the database
func (r *DatabaseRepository) AddCocktail(cocktail *model.Cocktail) {
	// Save the cocktail and its related instructions and ingredients
	result := r.db.Save(cocktail)
	if result.Error != nil {
		// Log the error or handle it appropriately
		// Since the interface doesn't allow returning an error, we'll just log it
		// In a real application, you might want to use a proper logging framework
		println("Error saving cocktail:", result.Error.Error())
	}
}
