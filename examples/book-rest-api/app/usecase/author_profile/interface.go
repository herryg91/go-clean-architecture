package author_profile_usecase

import "github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"

type UseCase interface {
	Get(id int) (*entity.AuthorProfile, error)
	Search(keyword string, sortBy SearchSortType, sortAsc bool) ([]*entity.AuthorProfile, error)
}

type SearchSortType string

const (
	SearchSortNoSort SearchSortType = "-"
	SearchSortAge    SearchSortType = "age"
	SearchSortName   SearchSortType = "name"
)
