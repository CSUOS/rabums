package v1

import (
	"fmt"
	"time"

	"github.com/LEE-WAN/RABUMS/cmd/database"
	"github.com/gin-gonic/gin"
)

//UserGet Healthcheck용 API
//절대로 프로그램적으로 사용하지 말것
func UserGet(c *gin.Context) {
	start := time.Now()
	db := database.Connect()
	defer db.Close()
	c.JSON(200, gin.H{
		"msg":   "Connect to db successfully! :-D",
		"takes": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	})
}
