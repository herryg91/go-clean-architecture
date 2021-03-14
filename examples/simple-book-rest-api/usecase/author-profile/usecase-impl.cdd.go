package author_profile

import (
	"sort"

	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) UseCase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) GetAuthorProfile(id int) (*entity.AuthorProfile, error) {
	return uc.repo.GetAuthorProfile(id)
}

func (uc *usecase) Search(keyword string, sortBy SearchSortType, sortAsc bool) ([]*entity.AuthorProfile, error) {
	result, err := uc.repo.Search(keyword)
	if err != nil {
		return []*entity.AuthorProfile{}, nil
	}

	switch sortBy {
	case SearchSortAge:
		sort.Slice(result, func(i, j int) bool {
			if sortAsc {
				return result[i].Age < result[j].Age
			}
			return result[i].Age > result[j].Age
		})
	case SearchSortName:
		sort.Slice(result, func(i, j int) bool {
			if sortAsc {
				return result[i].Name < result[j].Name
			}
			return result[i].Name > result[j].Name
		})
	}
	return result, err
}
