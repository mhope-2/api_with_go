package data

import (
	"github.com/mhope-2/api_with_go/internal/validator"
	// "github.com/lib/pq"
	"database/sql"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	ISBN      string    `json:"isbn"`
	Year      int32     `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `json:"omitempty"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	// title validation
	v.Check(book.Title != "", "title", "must not be empty")
	v.Check(len(book.Title) <= 100, "title", "must not be more than 100 characters")

	// year validation
	v.Check(book.Year != 0, "year", "must be a valid year")
	v.Check(book.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	// ISBN validation
	v.Check(book.ISBN != "", "isbn", "must not be empty")
	v.Check(len(book.ISBN) <= 5, "isbn", "must be more than 5 character")
}

type MockBookModel struct{}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

func (b MovieModel) Insert(book *Book) error {
	// Define the SQL query for inserting a new record in the movies table and returning
	// the system-generated data.
	query := `
	INSERT INTO book (title, isbn, year, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at`

	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.

	args := []interface{}{book.Title, book.ISBN, book.Year, book.CreatedAt}

	// Use the QueryRow() method to execute the SQL query on our connection pool,
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the movie struct.
	return b.DB.QueryRow(query, args...).Scan(&book.ID, &book.ISBN, &book.CreatedAt)
}

func (m MovieModel) Get(id int64) (*Book, error) {
// Mock the action...
	return nil,nil
}

func (m MovieModel) Update(Book *Book) error {
// Mock the action...
	return nil
}

func (m MovieModel) Delete(id int64) error {
// Mock the action...
	return nil
}