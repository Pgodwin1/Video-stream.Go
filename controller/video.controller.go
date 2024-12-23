package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) model.Video
}

type controller struct {
	service service.VideoService
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) model.Video {
	var video model.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
