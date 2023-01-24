package main

import (
	"api-pos/db"
	"github.com/gin-gonic/gin"
)

func main() {

	//connect db
	db.Connectdb()
	db.Migrate()

	r := gin.Default()
	r.Run()
}

