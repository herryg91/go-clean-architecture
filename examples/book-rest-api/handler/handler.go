package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/book_page"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}
type restHandler struct {
	book_page_uc book_page.UseCase
}

func NewHandler(book_page_uc book_page.UseCase) RestHandler {
	return &restHandler{book_page_uc: book_page_uc}
}

func (h *restHandler) GetBook(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.book_page_uc.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, book_page.ErrBookNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

type SearchBookReq struct {
	Keyword        string `json:"keyword"`
	SortBy         string `json:"sort"`
	SortDescending bool   `json:"descending"`
}

func (h *restHandler) SearchBook(c *gin.Context) {
	param := &SearchBookReq{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	data, err := h.book_page_uc.Search(param.Keyword, book_page.SearchSortType(param.SortBy), !param.SortDescending)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
