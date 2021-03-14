package handler

import (
	author_profile "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/author-profile"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/crud_authors"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/crud_books"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}
type restHandler struct {
	crudAuthorsUsecase   crud_authors.UseCase
	crudBooksUsecase     crud_books.UseCase
	authorProfileUsecase author_profile.UseCase
}

func NewHandler(crudAuthorsUsecase crud_authors.UseCase, crudBooksUsecase crud_books.UseCase, authorProfileUsecase author_profile.UseCase) RestHandler {
	return &restHandler{crudAuthorsUsecase: crudAuthorsUsecase, crudBooksUsecase: crudBooksUsecase, authorProfileUsecase: authorProfileUsecase}
}
