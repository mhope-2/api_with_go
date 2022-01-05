package data

import (
	"github.com/mhope-2/api_with_go/internal/validator"
	// "github.com/lib/pq" 
	"time"
	"database/sql"
)


type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	ISBN      string    `json:"isbn"`
	Year      int32     `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"omitempty"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	// title validation
	v.Check(book.Title != "", "title", "must not be empty")
	v.Check(len(book.Title) <= 100, "title", "must not be more than 100 characters")

	// year validation
	v.Check(book.Year !=0, "year", "must be a valid year")
	v.Check(book.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	// ISBN validation
	v.Check(book.ISBN != "", "isbn", "must not be empty")
	v.Check(len(book.ISBN) <= 5, "isbn", "must be more than 5 character")
}

// CRUD operation methods
// Define a BookModel struct type which wraps a sql.DB connection pool.
type BookModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the book table.
func (m BookModel) Insert(book *Book) error {
	return nil
}
// Add a placeholder method for fetching a specific record from the book table.
func (m BookModel) Get(id int64) (*Book, error) {
	return nil, nil
}

// Add a placeholder method for updating a specific record in the book table.
func (m BookModel) Update(book *Book) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the book table.
func (m BookModel) Delete(id int64) error {
	return nil
}