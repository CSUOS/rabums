package main

import (
	"github.com/gin-gonic/gin"

	"github.com/CSUOS/rabums/cmd/database"
	utils "github.com/CSUOS/rabums/cmd/utils"
	v1 "github.com/CSUOS/rabums/cmd/v1"
)

func main() {
	utils.LoadSettings()
	database.DBInit()
	utils.GenerateSecret()

	router := gin.Default()

	version1 := router.Group("/v1")
	{
		version1.GET("/user", v1.UserGet)
		version1.POST("/user", v1.UserPost)

		version1.PUT("/client", v1.ClientPut)
	}

	router.Run(":8080")
}
