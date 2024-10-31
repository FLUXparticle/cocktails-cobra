package model

type Cocktail struct {
	CocktailID   uint           `json:"id"`
	Name         string         `json:"name"`
	Instructions []*Instruction `json:"instructions"`
}
