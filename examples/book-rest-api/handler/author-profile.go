package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	author_profile_usecase "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/author_profile"
)

func (h *restHandler) GetAuthorProfile(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.author_profile_uc.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, author_profile_usecase.ErrProfileNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

type SearchAuthorReq struct {
	Keyword        string `json:"keyword"`
	SortBy         string `json:"sort"`
	SortDescending bool   `json:"descending"`
}

func (h *restHandler) SearchAuthor(c *gin.Context) {
	param := &SearchAuthorReq{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.author_profile_uc.Search(param.Keyword, author_profile_usecase.SearchSortType(param.SortBy), !param.SortDescending)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
