package models

import "os"
import "log"
import "encoding/json"

type Recipe struct {
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients,omitempty"`
	Instructions []string `json:"instructions,omitempty"`

	Servings        int    `json:"servings,omitempty"`
	PreparationTime string `json:"preparation_time,omitempty"`
	CookingTime     string `json:"cooking_time,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

func GetExample() (*Recipe, error) {

	var recipe Recipe

	logger := log.New(os.Stdout, "RECIPE: ", log.LstdFlags)
	readFile, err := os.Open("/Users/nfranzen/go/src/github.com/ntfrnzn/bakingdish/test/recipe.json")
	if err != nil {
		logger.Output(0, err.Error())
		return nil, err
	}

	jsonParser := json.NewDecoder(readFile)
	if err = jsonParser.Decode(&recipe); err != nil {
		logger.Output(0, "parsing recipe file "+err.Error())
		return nil, err
	}
	return &recipe, nil
}
