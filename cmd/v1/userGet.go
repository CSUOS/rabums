package v1

import (
	"fmt"
	"time"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/gin-gonic/gin"
)

//UserGet Healthcheck용 API
//절대로 프로그램적으로 사용하지 말것
func UserGet(c *gin.Context) {
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
