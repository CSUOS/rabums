package v1frontend

import (
	"net/http"
	"time"

	"github.com/CSUOS/rabums/pkg/database"
	"github.com/gin-gonic/gin"
)

type responseUserLog struct {
	ClientID int       `json:"clientId"`
	Event    int       `json:"event"`
	Date     time.Time `json:"date"`
}

//GetLogs 사용자의 로그 가져옴
func GetLogs(c *gin.Context) {
	userID, err := getUserID(c)

	if err != nil {
		panic(err)
	}

	user := database.UserInfo{}
	user.Get(userID)
	data, _ := user.GetLogs(0, 100)

	res := make([]responseUserLog, len(*data))
	for i := range *data {
		res[i].ClientID = (*data)[i].Client
		res[i].Event = int((*data)[i].Event)
		res[i].Date = (*data)[i].CreatedAt
	}
	c.JSON(http.StatusOK, res)
}
