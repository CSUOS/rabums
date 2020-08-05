package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/CSUOS/rabums/cmd/database"
	utils "github.com/CSUOS/rabums/cmd/utils"
	v1 "github.com/CSUOS/rabums/cmd/v1"
)

func main() {
	utils.LoadSettings()
	database.DBInit()
	utils.GenerateSecret()

	fmt.Println(utils.RabumsHASH(utils.RabumsHASH("test2")))
	fmt.Println(utils.RabumsHashedPasswdToken("962d3b4a8f231a9d9902619e1775648ee8db3ac90966ad013a27bdfa24940f93", "dcd08e8f5ef7dbbc06364709a595ad7269e78f8524df674477624c671708093c"))

	router := gin.Default()

	version1 := router.Group("/v1")
	{
		version1.GET("/user", v1.UserGet)
		version1.POST("/user", v1.UserPost)
		version1.PUT("/user/request/token", v1.UserReqTokenPut)
		version1.PUT("/user/request/register", v1.UserReqRegisterPut)

		version1.PUT("/client", v1.ClientPut)
	}

	router.Run(":8080")
}
