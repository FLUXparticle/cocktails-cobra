package server

import (
	"cocktails-cobra/pkg/cocktails/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RatingsHandler struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewRatingsHandler(log *zap.Logger, db *gorm.DB) *RatingsHandler {
	return &RatingsHandler{
		log: log,
		db:  db,
	}
}

func (h *RatingsHandler) AddRating(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Cocktail-ID"})
		return
	}

	var rating common.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating.CocktailID = uint(id)
	if err := h.db.Create(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Hinzufügen der Bewertung"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *RatingsHandler) GetAverageRating(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Cocktail-ID"})
		return
	}

	var average float64
	h.db.Model(&common.Rating{}).
		Where("cocktail_id = ?", id).
		Select("AVG(score)").
		Row().
		Scan(&average)

	c.JSON(http.StatusOK, gin.H{"average_rating": average})
}
