package author_profile

import "github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"

type SearchSortType string

const (
	SearchSortNoSort SearchSortType = "-"
	SearchSortAge    SearchSortType = "age"
	SearchSortName   SearchSortType = "name"
)

type Repository interface {
	GetAuthorProfile(id int) (*entity.AuthorProfile, error)
	Search(keyword string) ([]*entity.AuthorProfile, error)
}

type UseCase interface {
	Search(keyword string, sortBy SearchSortType, sortAsc bool) ([]*entity.AuthorProfile, error)
	GetAuthorProfile(id int) (*entity.AuthorProfile, error)
}
