package common

type Cocktail struct {
	ID           uint           `gorm:"primarykey"`
	Name         string         `json:"name"`
	Instructions []*Instruction `json:"instructions"`
}
