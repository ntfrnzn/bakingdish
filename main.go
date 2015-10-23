package main

import (
	"net/http"
	"github.com/ntfrnzn/bakingdish/routes"
)


func main() {

	router := routes.SetupTest()
	http.ListenAndServe("localhost:3000", router) // Start the server!

}
