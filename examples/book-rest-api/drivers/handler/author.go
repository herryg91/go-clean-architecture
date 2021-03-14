package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/usecase/crud_authors"
)

func (h *restHandler) GetAuthors(c *gin.Context) {
	data, err := h.crudAuthorsUsecase.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) GetAuthor(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.crudAuthorsUsecase.Get(id)
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

func (h *restHandler) CreateAuthor(c *gin.Context) {
	param := &entity.Author{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.crudAuthorsUsecase.Create(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) UpdateAuthor(c *gin.Context) {
	param := &entity.Author{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	paramId := c.Param("id")
	param.Id, _ = strconv.Atoi(paramId)

	result, err := h.crudAuthorsUsecase.Update(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
