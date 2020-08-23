package v1

import (
	"net/http"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/CSUOS/rabums/cmd/utils"

	"github.com/gin-gonic/gin"
)

//RequestClientUpdate /v1/client PUT으로 들어오는 요청
type RequestClientUpdate struct {
	MasterKey   string `json:"masterkey"`
	ClientID    string `json:"clientId" binding:"required"`
	ClientPW    string `json:"clientPw" binding:"required"`
	ChangedPw   string `json:"changedPw"`
	Link        string `json:"link" binding:"required"`
	Description string `json:"description"`
	Valid       bool   `json:"valid"`
	RenewToken  bool   `json:"renewToken"`
}

//ResponseClientUpdate /v1/client PUT 요청에 대한 응답
type ResponseClientUpdate struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
}

//ClientPut 클라이언트 생성 혹은 업데이트를 해주는 API
//->RequestClientUpdate
//<-ResponseClientUpdate
//400, 401, 500 에러 발생가능
func ClientPut(c *gin.Context) {
	req := RequestClientUpdate{
		Valid: false,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := database.GetClientInfo(req.ClientID)
	if err != nil {
		panic(err)
	}
	validMasterKey := utils.CheckIsMasterkey(req.MasterKey)
	isValid := validMasterKey ||
		client.ClientPW == req.ClientPW ||
		client.ClientPW == ""
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a valid key"})
		database.RecordLog(0, 0, database.INVALIDMASTERKEY, database.Message{
			"masterKey": req.MasterKey,
			"clientId":  req.ClientID,
			"clientPw":  req.ClientPW,
		})
		return
	}
	if !validMasterKey {
		req.Valid = client.Valid
	}

	client.ClientID = req.ClientID
	if req.MasterKey == "" {
		client.ClientPW = req.ClientPW
	}
	if req.ChangedPw != "" {
		client.ClientPW = req.ChangedPw
	}
	if req.Link != "" {
		client.Link = req.Link
	}
	if req.RenewToken || client.Token == "" {
		client.Token = utils.GenerateNewToken()
	}
	if req.Description != "" {
		client.Description = req.Description
	}
	client.Valid = req.Valid

	err = database.UpdateClientInfo(client)

	res := ResponseClientUpdate{
		Token: client.Token,
		Valid: client.Valid,
	}
	c.JSON(http.StatusOK, res)
	database.RecordLog(0, client.ID, database.UPDATED, database.Message{
		"clientId": client.ClientID,
	})
}
