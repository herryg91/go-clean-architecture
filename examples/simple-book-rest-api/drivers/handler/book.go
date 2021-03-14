package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/entity"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/crud_authors"
)

func (h *restHandler) GetBooks(c *gin.Context) {
	data, err := h.crudBooksUsecase.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) GetBook(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.crudBooksUsecase.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, crud_authors.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

type CreateBookReq struct {
	entity.Book
	Authors []int `json:"authors"`
}

func (h *restHandler) CreateBook(c *gin.Context) {
	param := &CreateBookReq{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	err = h.crudBooksUsecase.Create(param.Book, param.Authors)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: "success"})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

type UpdateBookReq struct {
	entity.Book
	Authors []int `json:"authors"`
}

func (h *restHandler) UpdateBook(c *gin.Context) {
	param := &UpdateBookReq{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	param.Id = id

	err = h.crudBooksUsecase.Update(param.Book, param.Authors)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: "success"})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
