package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hjkelly/discipline/config"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	// Load the config and bail if we encounter an error.
	configPtr, configErr := config.ParseConfig()
	if configErr != nil {
		log.Printf("Giving up because we couldn't load a valid config: %s", configErr.Error())
		return
	}

	router := httprouter.New()

	// Create the middleware handler
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	// After all middleware, go to the routes.
	n.UseHandler(router)

	// Start the server!
	http.ListenAndServe(":"+strconv.Itoa(configPtr.Port), n)
}
