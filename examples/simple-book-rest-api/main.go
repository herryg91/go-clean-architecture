package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/config"
	author_datasource "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/drivers/datasource/mysql/author"
	book_datasource "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/drivers/datasource/mysql/book"
	book_author_datasource "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/drivers/datasource/mysql/book_author"

	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/drivers/handler"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/pkg/mysql"
	author_profile "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/author-profile"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/crud_authors"
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/usecase/crud_books"
)

func main() {
	cfg := config.New()
	router := gin.Default()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPass, cfg.DBDatabaseName, cfg.DBLogMode)
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	authorDatasource := author_datasource.NewMysqlDatasource(db)
	bookDatasource := book_datasource.NewMysqlDatasource(db)
	bookAuthorDatasource := book_author_datasource.NewMysqlDatasource(db)

	crudAuthorsRepo := crud_authors.NewRepository(db, authorDatasource)
	crudAuthorsUsecase := crud_authors.NewUsecase(crudAuthorsRepo)

	crudBooksRepo := crud_books.NewRepository(db, bookDatasource, authorDatasource, bookAuthorDatasource)
	crudBooksUsecase := crud_books.NewUsecase(crudBooksRepo)

	authorProfileRepo := author_profile.NewRepository(db, bookDatasource, authorDatasource, bookAuthorDatasource)
	authorProfileUsecase := author_profile.NewUsecase(authorProfileRepo)

	h := handler.NewHandler(crudAuthorsUsecase, crudBooksUsecase, authorProfileUsecase)

	/* CRUD author for CMS */
	router.GET("/cms/author", h.GetAuthors)
	router.GET("/cms/author/:id", h.GetAuthor)
	router.POST("/cms/author", h.CreateAuthor)
	router.POST("/cms/author/:id", h.UpdateAuthor)

	/* CRUD book for CMS */
	router.GET("/cms/book", h.GetBooks)
	router.GET("/cms/book/:id", h.GetBook)
	router.POST("/cms/book", h.CreateBook)
	router.POST("/cms/book/:id", h.UpdateBook)

	router.GET("/author/:id/profile", h.GetAuthorProfile)
	router.POST("/search/author", h.SearchAuthor)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
