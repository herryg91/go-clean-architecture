package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetVideos(c *gin.Context)
	GetVideo(c *gin.Context)
	CreateVideo(c *gin.Context)
	UpdateVideo(c *gin.Context)
	DeleteVideo(c *gin.Context)
}
