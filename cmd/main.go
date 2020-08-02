package main

import (
	"github.com/gin-gonic/gin"

	"github.com/LEE-WAN/RABUMS/cmd/database"
	"github.com/LEE-WAN/RABUMS/cmd/util"
	v1 "github.com/LEE-WAN/RABUMS/cmd/v1"
)

func main() {
	LoadSettings()
	database.DBInit()
	util.GenerateSecret()

	router := gin.Default()

	version1 := router.Group("/v1")
	{
		version1.GET("/user", v1.UserGet)
		version1.POST("/user", v1.UserPost)

		version1.PUT("/client", v1.ClientPut)
	}

	router.Run(":8080")
}
