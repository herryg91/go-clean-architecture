package entity

import (
	"time"
)

type Book struct {
	Id           int        `json:"id"`
	Title        string     `json:"title"`
	ReleasedYear int        `json:"released_year"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
