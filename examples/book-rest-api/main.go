package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/app/usecase/book_page"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/config"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/handler"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/repository/book_repository_v1"
	"gorm.io/gorm/logger"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/mysql"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/redis"
)

func main() {
	cfg := config.New()
	router := gin.Default()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName, logger.LogLevel(cfg.DBLogMode))
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}
	rdsPool, err := redis.Connect(cfg.RedisHost, cfg.RedisPort, "")
	if err != nil {
		log.Panicln("Failed to Initialized redis:", err)
	}
	book_repo := book_repository_v1.New(db, rdsPool)
	book_page_uc := book_page.NewUseCase(book_repo)

	h := handler.NewHandler(book_page_uc)

	/* CRUD author for CMS */
	router.GET("/book/:id", h.GetBook)
	router.POST("/book/search", h.SearchBook)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
