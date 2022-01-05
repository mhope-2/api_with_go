package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) routes() *httprouter.Router {

	router := httprouter.New()

	// convert the notFoundResponse() helper to http.Handler using
	// the http.HandlerFunc() adapter and then set it as the
	// custom error handler for 404

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// convert the methodNotAllowedResponse() helper to a http.Handler
	// and set it as the custom handler for 405 Method Not Allowed Responses

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// register the relevant methods, url patterns and handler functions
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	// return the httprouter instance
	return router
}
