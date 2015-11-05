package models

import (
	"encoding/json"
	"github.com/gocraft/web"
	"os"
	"io/ioutil"
	"strings"
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
	readFile, err := os.Open( GetTestRecipeDir() + id + ".json")
	if err != nil {
		return nil, err
	}

	var recipe Recipe
	jsonParser := json.NewDecoder(readFile)
	if err = jsonParser.Decode(&recipe); err != nil {
		return nil, err
	}
	return &recipe, nil
}



func getAllIds() map[string]struct{} {

	var filenames map[string]struct{}
	filenames = make(map[string]struct{})

	contents, err := ioutil.ReadDir( GetTestRecipeDir() )
	if (err != nil){
		return filenames
	}	
	for _, value := range contents {
                fname := value.Name()
		if strings.HasSuffix(fname, ".json") {
			filenames[strings.TrimSuffix(fname, ".json")] = struct{}{}
			web.Logger.Println(" -- " + fname)
		}
	}
	return filenames
}


func (TestingApi) SearchRecipe(query *Query) ([]RecipeId, error) {
	web.Logger.Println(" -- search recipes for id <<" + query.Id + ">>")
	result := []RecipeId{}
        ids := getAllIds()
	if query.Id == "" {
		for id, _ := range ids {
			recipe, err := TestingApi{}.GetRecipe(id)
			if ( err == nil) {
				result = append(result, RecipeId{id,recipe.Name})
			}
		}
	} else {
		if _, ok := ids[query.Id]; ok {
			recipe, err := TestingApi{}.GetRecipe(query.Id)
			if ( err == nil) {
				result = append(result, RecipeId{query.Id,recipe.Name})
			}
		}
	}
	return result, nil
}

func (TestingApi) PostRecipe(recipe *Recipe) (string, error) {
	content, _ := json.Marshal(recipe)
	id := recipe.GetId()
	web.Logger.Println(id + "\n" +string(content))
	content, err := json.Marshal(recipe)
	if err != nil {
		return "", err
	}
	web.Logger.Println("posted the recipe: " + string(content))
	return "not_a_real_id", nil
}
