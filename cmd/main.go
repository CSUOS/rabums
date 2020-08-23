package main

import (
	"github.com/gin-gonic/contrib/static"
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

	router.Use(static.Serve("/", static.LocalFile("../dist", true)))

	version1 := router.Group("api/v1")
	{
		version1.GET("/ping", v1.DatabaseGet)

		version1.POST("/user", v1.UserPost)
		version1.PUT("/user/request/token", v1.UserReqTokenPut)
		version1.PUT("/user/request/register", v1.UserReqRegisterPut)

		version1.GET("/client", v1.ClientGet)
		version1.PUT("/client", v1.ClientPut)
	}

	router.Run(":8080")
}
