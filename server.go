package main

import (
	"fmt"
	"log"
	//	"github.com/corneldamian/json-binding"
	"encoding/json"
	"github.com/gocraft/web"
	"github.com/ntfrnzn/bakingdish/models"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func (c *Context) EchoParams(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, req.PathParams["param1"])

}

func (c *Context) TestRecipe(rw web.ResponseWriter, req *web.Request) {
	//r := models.Recipe{}
	//r.Name="Toasted Cheese"
	r, _ := models.GetExample()
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

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(recipe); err != nil {
		panic(err)
	}
	content, _ := json.Marshal(recipe)
	log.Println(string(content))
}

func main() {

	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
		// documentation for the below is not terribly clear -- in order to get the automatic loading of
		// index.html, it _must_ be specified as a web.StaticOption
		Middleware(web.StaticMiddlewareFromDir(http.Dir("static"), web.StaticOption{Prefix: "", IndexFile: "index.html"})).
		Post("/recipe", (*Context).PostRecipe).
		Get("/recipe", (*Context).TestRecipe).
		Get("/:param1", (*Context).EchoParams)
	http.ListenAndServe("localhost:3000", router) // Start the server!
}
