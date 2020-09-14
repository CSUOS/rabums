package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/CSUOS/rabums/pkg/database"
	utils "github.com/CSUOS/rabums/pkg/utils"
	v1 "github.com/CSUOS/rabums/pkg/v1"
	"github.com/CSUOS/rabums/pkg/v1frontend"
)

func main() {
	utils.LoadSettings()
	database.DBInit()
	utils.GenerateSecret()

	u := database.UserInfo{}
	u.Get("train96")
	tmp, _ := u.GetLogs(0, 10)
	fmt.Println(tmp)

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

	router.GET("/api/v1frontend/token", v1frontend.GetToken)
	router.POST("/api/v1frontend/login", v1frontend.LoginHandler)
	router.GET("/api/v1frontend/logout", v1frontend.LogoutHandler)
	version1frontend := router.Group("api/v1frontend")
	{
		version1frontend.Use(v1frontend.AuthMiddleware)
		version1frontend.GET("/logs", v1frontend.GetLogs)
	}

	router.Run(":8080")
}
