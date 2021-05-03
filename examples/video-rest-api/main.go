package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/app/usecase/crud_video"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/config"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/handler"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/repository/video_repository_v1"
	"gorm.io/gorm/logger"

	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/pkg/mysql"
)

func main() {
	cfg := config.New()
	router := gin.Default()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName, logger.LogLevel(cfg.DBLogMode))
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	video_repo := video_repository_v1.New(db)
	crud_video_uc := crud_video.NewUseCase(video_repo)

	h := handler.NewHandler(crud_video_uc)

	router.GET("/video", h.GetVideos)
	router.GET("/video/:id", h.GetVideo)
	router.POST("/video", h.CreateVideo)
	router.POST("/video/:id", h.UpdateVideo)
	router.DELETE("/video/:id", h.DeleteVideo)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
