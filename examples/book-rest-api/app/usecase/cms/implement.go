package cms_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
)

type usecase struct {
	book_repo   repository.BookRepository
	author_repo repository.AuthorRepository
}

func NewUseCase(book_repo repository.BookRepository, author_repo repository.AuthorRepository) UseCase {
	return &usecase{book_repo: book_repo, author_repo: author_repo}
}

func (uc *usecase) GetBook(id int) (*entity.Book, error) {
	data, err := uc.book_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	authorIds := []int{}
	for _, author := range data.Authors {
		authorIds = append(authorIds, author.Id)
	}
	mapOfAuthors, err := uc.author_repo.MultiGet(authorIds)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	for i, author := range data.Authors {
		if val, ok := mapOfAuthors[author.Id]; ok {
			data.Authors[i].Name = val.Name
		}
	}

	return data, nil

}
func (uc *usecase) GetBooks() ([]*entity.Book, error) {
	books, err := uc.book_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	authorIds := []int{}
	for _, book := range books {
		for _, author := range book.Authors {
			authorIds = append(authorIds, author.Id)
		}
	}
	mapOfAuthors, err := uc.author_repo.MultiGet(authorIds)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	for i, book := range books {
		for j, author := range book.Authors {
			if val, ok := mapOfAuthors[author.Id]; ok {
				books[i].Authors[j].Name = val.Name
			}
		}
	}
	return books, nil
}
func (uc *usecase) CreateBook(in entity.Book) (int, error) {
	data, err := uc.book_repo.Create(in)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data.Id, nil
}

func (uc *usecase) GetAuthor(id int) (*entity.Author, error) {
	data, err := uc.author_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrAuthorNotFound) {
			return nil, ErrAuthorNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
func (uc *usecase) GetAuthors() ([]*entity.Author, error) {
	data, err := uc.author_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil

}
func (uc *usecase) CreateAuthor(in entity.Author) (int, error) {
	data, err := uc.author_repo.Create(in)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data.Id, nil
}
