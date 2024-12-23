package main

import (
	"github.com/gin-gonic/gin"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"gilab.com/pragmaticreviews/golang-gin-poc/controller"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
       ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
       ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}