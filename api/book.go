package main

import (
	"fmt"
	"github.com/mhope-2/api_with_go/internal/database"
	"github.com/mhope-2/api_with_go/internal/validator"
	"net/http"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint.
func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	// Declare an anonymous struct to hold the information that we expect
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		// Use the new badRequestResponse() helper.
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Movie struct.
	movie := &database.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: database.Runtime(input.Runtime),
		Genres:  input.Genres,
	}

	// Initialize a new Validator.
	v := validator.New()

	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.

	if database.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}
