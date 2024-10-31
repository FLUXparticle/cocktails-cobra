package common

import "fmt"

type Instruction struct {
	ID         uint `gorm:"primarykey"`
	CocktailID uint
	CL         int    `json:"cl,omitempty"`
	Ingredient string `json:"ingredient"`
}

func (i *Instruction) String() string {
	if i.CL == 0 {
		return i.Ingredient
	} else {
		return fmt.Sprintf("%dcl %s", i.CL, i.Ingredient)
	}
}
