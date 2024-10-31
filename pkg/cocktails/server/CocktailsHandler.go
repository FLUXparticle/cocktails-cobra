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

func (h *CocktailsHandler) CocktailHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi!"))
}

func (h *CocktailsHandler) CocktailList(c *gin.Context) {
	var cocktails []common.Cocktail
	h.db.Find(&cocktails)

	result := make(map[uint]string)
	for _, cocktail := range cocktails {
		result[cocktail.ID] = cocktail.Name
	}

	time.Sleep(time.Second)

	c.JSON(200, result)
}

func (h *CocktailsHandler) CocktailDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var cocktail common.Cocktail
	h.db.Preload(clause.Associations).First(&cocktail, id)

	time.Sleep(time.Second)

	c.JSON(200, cocktail)
}
