package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mhope-2/api_with_go/internal/data"
	"github.com/mhope-2/api_with_go/internal/validator"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint.
func (app *Application) createBookHandler(w http.ResponseWriter, r *http.Request) {

	// Declare an anonymous struct to hold the information that we expect
	var input struct {
		Title     string    `json:"title"`
		ISBN      string    `json:"isbn"`
		Year      int32     `json:"year"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		// Use the new badRequestResponse() helper.
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Book struct.
	book := &data.Book{
		Title:     input.Title,
		ISBN:      input.ISBN,
		Year:      input.Year,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Initialize a new Validator.
	v := validator.New()

	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.

	if data.ValidateBook(v, book); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// insert book record into model
	err = app.data.Movies.Insert(book)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// When sending a HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header,
	// interpolating the system-generated ID for our new movie in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", book.ID))
	// Write a JSON response with a 201 Created status code, the movie data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": book}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)

	}

}
