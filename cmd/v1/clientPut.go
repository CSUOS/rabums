package v1

import (
	"net/http"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/CSUOS/rabums/cmd/utils"

	"github.com/gin-gonic/gin"
)

//RequestClientUpdate /v1/client PUT으로 들어오는 요청
type RequestClientUpdate struct {
	MasterKey string `json:"masterkey" binding:"required"`
	ClientID  string `json:"clientId" binding:"required"`
	ClientPW  string `json:"clientPw" binding:"required"`
	Link      string `json:"link" binding:"required"`
	Valid     bool   `json:"valid"`
}

//ResponseClientUpdate /v1/client PUT 요청에 대한 응답
type ResponseClientUpdate struct {
	Token string `json:"token"`
}

//ClientPut 클라이언트 생성 혹은 업데이트를 해주는 API
//->RequestClientUpdate
//<-ResponseClientUpdate
//400, 401, 500 에러 발생가능
func ClientPut(c *gin.Context) {
	req := RequestClientUpdate{
		Valid: true,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isValid := utils.CheckIsMasterkey(req.MasterKey); isValid != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a valid master key"})
		database.RecordLog(0, 0, database.INVALIDMASTERKEY, database.Message{
			"masterKey": req.MasterKey,
			"clientId":  req.ClientID,
			"clientPw":  req.ClientPW,
		})
		return
	}

	client, err := database.GetClientInfo(req.ClientID)
	if err != nil {
		panic(err)
	}

	client.ClientID = req.ClientID
	client.ClientPW = req.ClientPW
	client.Link = req.Link
	client.Token = utils.GenerateNewToken()
	client.Valid = req.Valid

	err = database.UpdateClientInfo(client)

	res := ResponseClientUpdate{
		Token: client.Token,
	}
	c.JSON(http.StatusOK, res)
	database.RecordLog(0, client.ID, database.UPDATED, database.Message{
		"clientId": client.ClientID,
	})
}
