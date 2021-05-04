package book_page

import "github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"

type UseCase interface {
	Get(id int) (*entity.Book, error)
	Search(keyword string, sortBy SearchSortType, sortAsc bool) ([]*entity.Book, error)
}

type SearchSortType string

const (
	SearchSortNoSort SearchSortType = "-"
	SearchSortYear   SearchSortType = "year"
	SearchSortName   SearchSortType = "title"
	SearchSortRating SearchSortType = "rating"
)
