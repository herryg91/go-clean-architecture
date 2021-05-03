package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/app/usecase/crud_video"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
)

type restHandler struct {
	crud_video_uc crud_video.UseCase
}

func NewHandler(crud_video_uc crud_video.UseCase) RestHandler {
	return &restHandler{crud_video_uc: crud_video_uc}
}

func (h *restHandler) GetVideos(c *gin.Context) {
	data, err := h.crud_video_uc.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) GetVideo(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.crud_video_uc.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, crud_video.ErrVideoNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

func (h *restHandler) CreateVideo(c *gin.Context) {
	param := &entity.Video{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.crud_video_uc.Create(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) UpdateVideo(c *gin.Context) {
	param := &entity.Video{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	paramId := c.Param("id")
	param.Id, _ = strconv.Atoi(paramId)

	result, err := h.crud_video_uc.Update(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) DeleteVideo(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.crud_video_uc.Delete(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: fmt.Sprintf("id:%d. successfully deleted", id)})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
