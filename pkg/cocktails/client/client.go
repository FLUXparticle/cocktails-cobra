package client

import (
	"cocktails-cobra/pkg/cocktails/common"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
)

func getCocktails() map[string]interface{} {
	response, err := http.Get("http://localhost:8080/cocktails")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	cocktails := make(map[string]interface{})
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &cocktails)
	if err != nil {
		panic(err)
	}

	return cocktails
}

func getCocktail(id string) *common.Cocktail {
	response, err := http.Get("http://localhost:8080/cocktails/" + id)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var cocktail *common.Cocktail

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &cocktail)
	if err != nil {
		panic(err)
	}

	return cocktail
}

func filterMilkIDs(cocktails map[string]interface{}) []string {
	milkIDs := make([]string, 0)
	for id, cocktail := range cocktails {
		if name, ok := cocktail.(string); ok {
			if strings.Contains(name, "Milk") {
				milkIDs = append(milkIDs, id)
			}
		}
	}

	//fmt.Println(milkIDs)

	return milkIDs
}

func DoSequential() int {
	cocktails := getCocktails()

	milkIDs := filterMilkIDs(cocktails)

	sumMilk := 0

	for _, id := range milkIDs {
		cocktail := getCocktail(id)
		//fmt.Printf("%s:\n", cocktail.Name)
		for _, instruction := range cocktail.Instructions {
			if instruction.Ingredient == "Milch" {
				sumMilk += instruction.CL
			}
			//fmt.Printf("  %s\n", instruction.String())
		}
	}
	return sumMilk
}

func DoParallel() int {
	cocktails := getCocktails()

	milkIDs := filterMilkIDs(cocktails)

	ch := make(chan int)

	wg := sync.WaitGroup{}

	for _, loopID := range milkIDs {
		id := loopID
		wg.Add(1)
		go func() {
			cocktail := getCocktail(id)
			for _, instruction := range cocktail.Instructions {
				if instruction.Ingredient == "Milch" {
					ch <- instruction.CL
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sumMilk := 0
	for milk := range ch {
		sumMilk += milk
	}
	return sumMilk
}
