package main

import (
	"api-pos/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//connect db
	db.Connectdb()
	db.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r := gin.Default()
	r.Use(cors.New(corsConfig))
	serveRoutes(r)
	r.Run(":8080")
}

