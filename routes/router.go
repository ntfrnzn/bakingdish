package routes

import (
	"encoding/json"
	"github.com/gocraft/web"
	"github.com/ntfrnzn/bakingdish/models"
	"io"
	"io/ioutil"
	// "reflect"
	// "log"
	"fmt"
	"net/http"
	// "reflect"
	// "github.com/corneldamian/json-binding"
	// "strings"
	// "github.com/gocraft/health"
)

type Context struct {
	api models.RecipeApi
}

func (c *Context) TestApiMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.api = models.TestingApi{}
	next(rw, req)
}

func handleUnmarshalError(rw web.ResponseWriter, err error) {
	fmt.Println(err.Error())
	http.Error(rw, err.Error(), http.StatusBadRequest)
	return
}

func getPostedBody(req *web.Request) ([]byte, error) {

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		return nil, err
	}
	if err := req.Body.Close(); err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	return body, nil
}

func (c *Context) GetRecipe(rw web.ResponseWriter, req *web.Request) {
	id := req.PathParams["id"]
	r, _ := (*c).api.GetRecipe(id)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(r)
}

func (c *Context) PostRecipe(rw web.ResponseWriter, req *web.Request) {
	var recipe models.Recipe

	body, err := getPostedBody(req)
	if err != nil {
		panic(err)
		return
	}

	err = json.Unmarshal(body, &recipe)
	if err != nil {
		handleUnmarshalError(rw, err)
		return
	}

	id, err := (*c).api.PostRecipe(&recipe)
	if err != nil {
		panic(err)
		return
	}

	var result models.RecipeId
	result.Id = id
	result.Name = recipe.Name
	
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(rw).Encode(result)

}

func (c *Context) SearchRecipe(rw web.ResponseWriter, req *web.Request) {

	body, err := getPostedBody(req)
	if err != nil {
		panic(err)
		return
	}

	var query models.Query
	err = json.Unmarshal(body, &query)
	if err != nil {
		handleUnmarshalError(rw, err)
		return
	}

	results, err := (*c).api.SearchRecipe(&query)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		panic(err)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(results) // prefaces with {"Offset": 0} ??

	}
}

func SetupTest() *web.Router {
	// regarding the static middleware, documentation for the
	// below is not terribly clear -- in order to get the
	// automatic loading of "/path/index.html" when requesting
	// just "/path/", it _must_ be specified as a web.StaticOption

	// Create your router
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).         // Use some included middleware
		Middleware(web.ShowErrorsMiddleware).     // ...
		Middleware((*Context).TestApiMiddleware). // use the Test version of the methods
		Middleware(web.StaticMiddlewareFromDir(http.Dir("static"), web.StaticOption{Prefix: "", IndexFile: "index.html"})).
		Post("/recipe", (*Context).PostRecipe).
		Get("/recipe/:id", (*Context).GetRecipe).
		Post("/search/", (*Context).SearchRecipe)
	return router
}
