package v1

import (
	"net/http"

	"github.com/LEE-WAN/RABUMS/cmd/database"
	"github.com/LEE-WAN/RABUMS/cmd/util"

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

	db := database.Connect()
	defer db.Close()

	if isValid := util.CheckIsMasterkey(req.MasterKey); isValid != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a valid master key"})
		return
	}

	var client database.ClientInfo
	isUpdate := true
	if err := db.Where(
		database.ClientInfo{ClientID: req.ClientID},
	).Take(&client).Error; err != nil {
		// 새로 생성하는 케이스
		isUpdate = false
	}
	client.ClientID = req.ClientID
	client.ClientPW = req.ClientPW
	client.Link = req.Link
	client.Token = util.GenerateNewToken()
	client.Valid = req.Valid
	if isUpdate {
		db.Save(&client)
	} else {
		db.Create(&client)
	}
	res := ResponseClientUpdate{
		Token: client.Token,
	}
	c.JSON(http.StatusOK, res)

}
