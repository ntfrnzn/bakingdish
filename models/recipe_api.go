package models

import (
	"encoding/json"
	"github.com/gocraft/web"
)

var testExample *Recipe

func init() {
	testExample, _ = GetExample()
}

type Query struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RecipeId struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RecipeApi interface {
	GetRecipe(id string) (*Recipe, error)
	PostRecipe(recipe *Recipe) (string, error)
	SearchRecipe(query *Query) ([]RecipeId, error)
}

type TestingApi struct {
}

func (TestingApi) GetRecipe(id string) (*Recipe, error) {
	web.Logger.Println("get recipe for id: " + id)
	return GetExample()
}

func (TestingApi) SearchRecipe(query *Query) ([]RecipeId, error) {
	web.Logger.Println(" -- search recipes for id <<" + query.Id + ">>")
	if query.Id == "" || query.Id == "dummy_id" {
		return []RecipeId{{"dummy_id", testExample.Name}}, nil
	}
	return []RecipeId{}, nil
}

func (TestingApi) PostRecipe(recipe *Recipe) (string, error) {
	content, _ := json.Marshal(recipe)
	web.Logger.Println(string(content))
	content, err := json.Marshal(recipe)
	if err != nil {
		return "", err
	}
	web.Logger.Println("posted the recipe: " + string(content))
	return "dummy_id", nil
}
