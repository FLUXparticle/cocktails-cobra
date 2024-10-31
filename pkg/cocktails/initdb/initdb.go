package initdb

import (
	"bufio"
	"cocktails-cobra/pkg/cocktails/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"strings"
)

func readCocktails(filename string) []*common.Cocktail {
	var cocktails []*common.Cocktail

	in, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var instructions []*common.Instruction
	var name string
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := sc.Text()
		if len(name) == 0 {
			name = line
		} else if len(line) > 0 {
			split := strings.Split(line, "cl:")
			switch len(split) {
			case 1:
				instructions = append(instructions, &common.Instruction{
					Ingredient: split[0],
				})
			case 2:
				if cl, err := strconv.Atoi(split[0]); err != nil {
					panic(err)
				} else {
					instructions = append(instructions, &common.Instruction{
						CL:         cl,
						Ingredient: split[1],
					})
				}
			}
		} else {
			cocktails = append(cocktails, &common.Cocktail{
				Name:         name,
				Instructions: instructions,
			})
			name = ""
			instructions = nil
		}
	}

	return cocktails
}

func writeCocktails(filename string, cocktails []*common.Cocktail) {
	os.Remove(filename)
	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(sqlite.Open(filename), cfg)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&common.Cocktail{})
	db.AutoMigrate(&common.Instruction{})

	for _, cocktail := range cocktails {
		db.Create(cocktail)
	}
}

func InitDB() {
	cocktails := readCocktails("cocktails.txt")
	writeCocktails("cocktails.db", cocktails)
}
