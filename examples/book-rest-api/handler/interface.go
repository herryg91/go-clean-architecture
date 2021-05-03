package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	// CRUD Author (for CMS)
	GetAuthors(c *gin.Context)
	GetAuthor(c *gin.Context)
	CreateAuthor(c *gin.Context)

	// CRUD Book (for CMS)
	GetBooks(c *gin.Context)
	GetBook(c *gin.Context)
	CreateBook(c *gin.Context)

	// Author Profile
	GetAuthorProfile(c *gin.Context)
	SearchAuthor(c *gin.Context)
}
