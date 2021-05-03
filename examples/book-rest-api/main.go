package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	author_profile_usecase "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/author_profile"
	cms_usecase "github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/cms"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/config"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/handler"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/repository/author_repository_v1"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/repository/book_repository_v1"
	"gorm.io/gorm/logger"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/mysql"
)

func main() {
	cfg := config.New()
	router := gin.Default()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName, logger.LogLevel(cfg.DBLogMode))
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	author_repo := author_repository_v1.New(db)
	book_repo := book_repository_v1.New(db)

	cms_uc := cms_usecase.NewUseCase(book_repo, author_repo)
	author_profile_uc := author_profile_usecase.NewUseCase(author_repo)

	h := handler.NewHandler(cms_uc, author_profile_uc)

	/* CRUD author for CMS */
	router.GET("/cms/author", h.GetAuthors)
	router.GET("/cms/author/:id", h.GetAuthor)
	router.POST("/cms/author", h.CreateAuthor)

	/* CRUD book for CMS */
	router.GET("/cms/book", h.GetBooks)
	router.GET("/cms/book/:id", h.GetBook)
	router.POST("/cms/book", h.CreateBook)

	router.GET("/author/:id/profile", h.GetAuthorProfile)
	router.POST("/search/author", h.SearchAuthor)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
