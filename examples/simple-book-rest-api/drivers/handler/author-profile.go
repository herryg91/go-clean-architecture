package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	author_profile "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/author-profile"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/crud_authors"
)

func (h *restHandler) GetAuthorProfile(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.authorProfileUsecase.GetAuthorProfile(id)
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

	result, err := h.authorProfileUsecase.Search(param.Keyword, author_profile.SearchSortType(param.SortBy), !param.SortDescending)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
