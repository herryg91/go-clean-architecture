package handler

import (
	author_profile_usecase "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/author_profile"
	cms_usecase "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/cms"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}
type restHandler struct {
	cms_uc            cms_usecase.UseCase
	author_profile_uc author_profile_usecase.UseCase
}

func NewHandler(cms_uc cms_usecase.UseCase, author_profile_uc author_profile_usecase.UseCase) RestHandler {
	return &restHandler{cms_uc: cms_uc, author_profile_uc: author_profile_uc}
}
