package model

import "fmt"

type Instruction struct {
	AmountCL   int         `json:"amount"`
	Ingredient *Ingredient `json:"ingredient"`
}

func (i *Instruction) String() string {
	if i.AmountCL == 0 {
		return i.Ingredient.Name
	} else {
		return fmt.Sprintf("%dcl %s", i.AmountCL, i.Ingredient.Name)
	}
}
