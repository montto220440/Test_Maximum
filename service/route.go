package main

import (
	"api-pos/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	Controller := controller.List{}
	Group := r.Group("/list")
	Group.GET("", Controller.FindAll)
	Group.GET("/:id", Controller.FindOne)
	Group.POST("", Controller.Create)
	Group.PATCH("/:id", Controller.Update)
	Group.DELETE("/:id", Controller.Delete)

}
