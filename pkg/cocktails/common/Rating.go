package common

type Rating struct {
	ID         uint `gorm:"primarykey"`
	CocktailID uint
	Score      int    `json:"score"`
	Review     string `json:"review,omitempty"`
}
