package author_profile_usecase

import (
	"errors"
	"fmt"
	"sort"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
)

type usecase struct {
	author_repo repository.AuthorRepository
}

func NewUseCase(author_repo repository.AuthorRepository) UseCase {
	return &usecase{author_repo: author_repo}
}

func (uc *usecase) Get(id int) (*entity.AuthorProfile, error) {
	author, err := uc.author_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrAuthorNotFound) {
			return nil, ErrProfileNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return author.ToProfile(), nil
}

func (uc *usecase) Search(keyword string, sortBy SearchSortType, sortAsc bool) ([]*entity.AuthorProfile, error) {
	authors, err := uc.author_repo.Search(keyword)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	profiles := []*entity.AuthorProfile{}
	for _, author := range authors {
		profiles = append(profiles, author.ToProfile())
	}

	switch sortBy {
	case SearchSortAge:
		sort.Slice(profiles, func(i, j int) bool {
			if sortAsc {
				return profiles[i].Age < profiles[j].Age
			}
			return profiles[i].Age > profiles[j].Age
		})
	case SearchSortName:
		sort.Slice(profiles, func(i, j int) bool {
			if sortAsc {
				return profiles[i].Name < profiles[j].Name
			}
			return profiles[i].Name > profiles[j].Name
		})
	}
	return profiles, nil
}
