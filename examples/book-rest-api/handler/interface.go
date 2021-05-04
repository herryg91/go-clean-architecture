package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetBook(c *gin.Context)
	SearchBook(c *gin.Context)
}
