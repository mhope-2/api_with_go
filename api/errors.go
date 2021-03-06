package main

import (
	"fmt"
	"net/http"
)

// function to log errors
func (app *Application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// error response
func (app *Application) errorResponse(w http.ResponseWriter, r *http.Request,
	status int, message interface{}) {

	env := envelope{"error": message}

	// write the response using the writeJSON helper funtion
	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
	}
}

// server error response
func (app *Application) serverErrorResponse(w http.ResponseWriter, r *http.Request,
	err error) {

	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"

	app.errorResponse(w, r, http.StatusNotFound, message)
}

// not found response
func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"

	app.errorResponse(w, r, http.StatusNotFound, message)
}

// menthod not allowed response
func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// bad request response
func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// failed validation response
func (app *Application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
