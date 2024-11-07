package server

import (
	"cocktails-cobra/pkg/cocktails/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"time"
)

type CocktailsHandler struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewCocktailsHandler(log *zap.Logger, db *gorm.DB) *CocktailsHandler {
	return &CocktailsHandler{
		log: log,
		db:  db,
	}
}

func (h *CocktailsHandler) CocktailList(c *gin.Context) {
	ingredient := c.Query("ingredient")

	var cocktails []common.Cocktail

	if ingredient != "" {
		// Cocktails suchen, die die angegebene Zutat enthalten
		h.db.Joins("JOIN instructions ON instructions.cocktail_id = cocktails.id").
			Where("instructions.ingredient = ?", ingredient).
			Group("cocktails.id").
			Find(&cocktails)
	} else {
		h.db.Find(&cocktails)
	}

	result := make(map[uint]string)
	for _, cocktail := range cocktails {
		result[cocktail.ID] = cocktail.Name
	}

	c.JSON(http.StatusOK, result)
}

func (h *CocktailsHandler) CocktailDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var cocktail common.Cocktail
	h.db.Preload(clause.Associations).First(&cocktail, id)

	time.Sleep(time.Second)

	c.JSON(http.StatusOK, cocktail)
}

// TODO weitere Handler Functions hier
func (h *CocktailsHandler) IngredientsList(c *gin.Context) {
	ingredientsMap := make(map[string]bool)

	var instructions []common.Instruction
	h.db.Find(&instructions)
	for _, instruction := range instructions {
		ingredientsMap[instruction.Ingredient] = true
	}

	var ingredients []string
	for ingredient := range ingredientsMap {
		ingredients = append(ingredients, ingredient)
	}

	c.JSON(http.StatusOK, ingredients)
}
