package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/CSUOS/rabums/pkg/database"
	utils "github.com/CSUOS/rabums/pkg/utils"
	v1 "github.com/CSUOS/rabums/pkg/v1"
)

func main() {
	utils.LoadSettings()
	database.DBInit()
	utils.GenerateSecret()

	c := database.ClientInfo{}
	if err := c.Get("rabums"); err != nil {
		c = database.ClientInfo{
			ClientID:    "rabums",
			ClientPW:    "nil",
			Link:        "https://rabums.csuos.ml",
			Description: "Rabums에서 자동으로 생성한 기본 계정입니다.",
			Valid:       true,
			Token:       utils.GenerateNewToken(),
		}
		err = c.Create()
		if err != nil {
			panic("DB Problem")
		}
	}

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

	router.Run(":8080")
}
