package main

import (
	"reflect"
	"encoding/json"
	"github.com/gocraft/web"
	"github.com/ntfrnzn/bakingdish/models"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
        "fmt"
	// "github.com/corneldamian/json-binding"
	// "strings"
        // "github.com/gocraft/health"
)


type Query struct {
}

type RecipeId struct {
	Id string
	Name string
}

type RecipeApi interface {
	GetRecipe(id string) (*models.Recipe, error)
	PostRecipe(recipe *models.Recipe) (string, error)
        SearchRecipe(query *Query) ([]RecipeId, error)
}


type TestingApi struct {
}


func (TestingApi) GetRecipe(id string) (*models.Recipe, error) {
	web.Logger.Println("get recipe for id: " + id)
	return models.GetExample() 
}

func (TestingApi) SearchRecipe(query *Query) ([]RecipeId, error) {
	web.Logger.Println("search recipes for id")
	return []RecipeId{ {"dummy_id","Indian-Style Grilled Tuna Steaks with Aromatic Spice Paste"} }, nil
}

func (TestingApi) PostRecipe( recipe *models.Recipe) (string, error) {
	content, _ := json.Marshal(recipe)
	web.Logger.Println(string(content))
	content, err := json.Marshal(recipe)
	if (err != nil){
		return "", err
	}
	web.Logger.Println( "posted the recipe: " + string(content))
        return "dummy_id", nil
}



type Context struct {
   api RecipeApi
}


func (c *Context) TestApiMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
        web.Logger.Println("here we are in middleware")
	c.api = TestingApi{}
	web.Logger.Println( reflect.TypeOf( c ).String() )
	web.Logger.Println( reflect.TypeOf( (*c).api ).String() )
        next(rw, req)
}




func (c *Context) GetRecipe(rw web.ResponseWriter, req *web.Request) {
        id := req.PathParams["id"]
        r, _ := (*c).api.GetRecipe(id)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(r)
}

func (c *Context) PostRecipe(rw web.ResponseWriter, req *web.Request) {
	var recipe models.Recipe
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &recipe); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}
	}

        id, err := (*c).api.PostRecipe(&recipe)
        if (err != nil ){
		panic(err)
	}
        
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, `{"id":"`+id+`"}"`)
}


func (c *Context) SearchRecipe(rw web.ResponseWriter, req *web.Request) {
	var query Query
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &query); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}
	}

        results, err := (*c).api.SearchRecipe(&query)
        if (err != nil ){
		panic(err)
	}
        
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(rw).Encode(results)
}



func main() {

	// regarding the static middleware (angular etc) documentation
	// for the below is not terribly clear -- in order to get the
	// automatic loading of index.html, it _must_ be specified as
	// a web.StaticOption

	// Create your router
	router := web.New( Context{}  ).
		Middleware(web.LoggerMiddleware).     // Use some included middleware
		Middleware(web.ShowErrorsMiddleware). // ...
		Middleware((*Context).TestApiMiddleware).   // use the Test version of the methods
		Middleware(web.StaticMiddlewareFromDir(http.Dir("static"), web.StaticOption{Prefix: "", IndexFile: "index.html"})).
		Post("/recipe", (*Context).PostRecipe).
		Get("/recipe/:id", (*Context).GetRecipe).
		Post("/search/", (*Context).SearchRecipe)
	http.ListenAndServe("localhost:3000", router) // Start the server!
}
