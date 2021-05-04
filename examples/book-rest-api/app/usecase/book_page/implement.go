package book_page

import (
	"errors"
	"fmt"
	"sort"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
)

type usecase struct {
	book_repo repository.BookRepository
}

func NewUseCase(book_repo repository.BookRepository) UseCase {
	return &usecase{book_repo: book_repo}
}

func (uc *usecase) Get(id int) (*entity.Book, error) {
	b, err := uc.book_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return b, nil
}

func (uc *usecase) Search(keyword string, sortBy SearchSortType, sortAsc bool) ([]*entity.Book, error) {
	books, err := uc.book_repo.Search(keyword)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	switch sortBy {
	case SearchSortYear:
		sort.Slice(books, func(i, j int) bool {
			if sortAsc {
				return books[i].ReleasedYear < books[j].ReleasedYear
			}
			return books[i].ReleasedYear > books[j].ReleasedYear
		})
	case SearchSortName:
		sort.Slice(books, func(i, j int) bool {
			if sortAsc {
				return books[i].Title < books[j].Title
			}
			return books[i].Title > books[j].Title
		})
	case SearchSortRating:
		sort.Slice(books, func(i, j int) bool {
			if sortAsc {
				return books[i].Rating < books[j].Rating
			}
			return books[i].Rating > books[j].Rating
		})
	}

	return books, nil
}
