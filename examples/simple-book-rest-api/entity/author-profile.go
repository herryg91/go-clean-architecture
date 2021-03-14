package entity

import (
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/pkg/helpers"
)

type AuthorProfile struct {
	Author
	Age   int    `json:"age"`
	Books []Book `json:"books"`
}

func (AuthorProfile) New(a Author, books []Book) *AuthorProfile {
	return &AuthorProfile{
		Author: a,
		Age:    helpers.CountAge(a.Birthdate.Time()),
		Books:  books,
	}
}
