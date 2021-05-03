package entity

import (
	"time"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/helpers"
)

type Author struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Birthdate helpers.Date `json:"birthdate"`
	CreatedAt *time.Time   `json:"created_at,omitempty"`
	UpdatedAt *time.Time   `json:"updated_at,omitempty"`
}

func (author *Author) ToBookAuthor() *BookAuthor {
	return &BookAuthor{
		Id:   author.Id,
		Name: author.Name,
	}
}

func (author *Author) ToProfile() *AuthorProfile {
	return &AuthorProfile{
		Author: *author,
		Age:    helpers.CountAge(author.Birthdate.Time()),
	}
}
