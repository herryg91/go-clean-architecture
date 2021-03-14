package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/config"
	video_datasource "github.com/herryg91/go-clean-architecture/examples/video-rest-api/drivers/datasource/mysql/videos"

	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/drivers/handler"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/pkg/mysql"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/usecase/crud_videos"
)

func main() {
	cfg := config.New()
	router := gin.Default()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName, cfg.DBLogMode)
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	videoDatasource := video_datasource.NewMysqlDatasource(db)

	crudVideosRepo := crud_videos.NewRepository(db, videoDatasource)
	crudVideosUsecase := crud_videos.NewUsecase(crudVideosRepo)

	h := handler.NewHandler(crudVideosUsecase)

	router.GET("/video", h.GetVideos)
	router.GET("/video/:id", h.GetVideo)
	router.POST("/video", h.CreateVideo)
	router.POST("/video/:id", h.UpdateVideo)
	router.DELETE("/video/:id", h.DeleteVideo)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
