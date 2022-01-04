package database

import (
	"github.com/mhope-2/internal/validator"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	ISBN      string    `json:"isbn"`
	Year      int32     `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	// title validation
	v.Check(book.Title != "", "title", "must not be empty")
	v.Check(len(book.Title) <= 100, "title", "must not be more than 100 characters")

	// year validation
	v.Check(book.year !=0, "year", "must be a valid year")
	v.Check(book.year <= int32(time.Now().Year()), "year", "must not be in the future")
}
