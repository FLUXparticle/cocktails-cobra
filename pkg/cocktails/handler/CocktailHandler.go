package handler

import (
	"cocktails-cobra/pkg/cocktails/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CocktailHandler handles HTTP requests for cocktails
type CocktailHandler struct {
	repo repository.MutableCocktailRepository
}

// NewCocktailHandler creates a new instance of CocktailHandler
func NewCocktailHandler(repo repository.MutableCocktailRepository) *CocktailHandler {
	return &CocktailHandler{
		repo: repo,
	}
}

// GetCocktails handles GET requests for all cocktails
func (h *CocktailHandler) GetCocktails(c *gin.Context) {
	cocktails, err := h.repo.GetCocktails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cocktails)
}

// RegisterRoutes registers the routes for the CocktailHandler
func (h *CocktailHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/cocktails", h.GetCocktails)
}
