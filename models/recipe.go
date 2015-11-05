package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
)

type Recipe struct {
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients,omitempty"`
	Instructions []string `json:"instructions,omitempty"`

	Servings        int    `json:"servings,omitempty"`
	PreparationTime string `json:"preparation_time,omitempty"`
	CookingTime     string `json:"cooking_time,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

func GetId(text string) string {
	// for two lines, this is actually a little tricky (for me?),
	// with conversions between string and []byte and the
	// details of slice <--> array exchange
	encode := md5.Sum([]byte(text))
	return hex.EncodeToString(encode[:])
}

// same as the previous, but now as a struct method
func (recipe *Recipe) GetId() string {
	if recipe.Name == "" {
		return "null_id"
	}
	recode := md5.Sum([]byte(recipe.Name))
	return hex.EncodeToString(recode[:])
}


func GetTestRecipeDir () string {
	return "/Users/nfranzen/go/src/github.com/ntfrnzn/bakingdish/test/recipes/"
}


func GetExample() (*Recipe, error) {
	var recipe Recipe

	logger := log.New(os.Stdout, "RECIPE: ", log.LstdFlags)
        //testDir := "/Users/nfranzen/go/src/github.com/ntfrnzn/bakingdish/test/recipes/"
	readFile, err := os.Open( GetTestRecipeDir() + "4a6168e3f784031f6b304319cd20101f.json")
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
