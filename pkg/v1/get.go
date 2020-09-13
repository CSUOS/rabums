package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/CSUOS/rabums/pkg/database"
	"github.com/gin-gonic/gin"
)

//DatabaseGet Healthcheck용 API
//절대로 프로그램적으로 사용하지 말것
func DatabaseGet(c *gin.Context) {
	start := time.Now()
	err := database.Ping()
	if err != nil {
		c.JSON(500, gin.H{
			"msg":   "Fail to connect to db. :-(",
			"takes": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg":   "Connect to db successfully! :-D",
			"takes": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		})
	}
}

//ClientGet Healthcheck용 API
//절대로 프로그램적으로 사용하지 말것
func ClientGet(c *gin.Context) {
	skip, _ := strconv.Atoi(c.Query("skip"))
	clients, err := database.GetClientList(skip)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":   "Fail to connect to db. :-(",
			"error": err.Error(),
		})
	}
	for _, client := range clients {
		client.ClientPW = ""
		client.Token = ""
	}

	c.JSON(http.StatusOK, clients)
}
