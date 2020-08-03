package v1

import (
	"net/http"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/CSUOS/rabums/cmd/utils"
	"github.com/gin-gonic/gin"
)

//RequestUserInfo /v1/user로 들어오는 POST 요청
type RequestUserInfo struct {
	Token  string `json:"token" binding:"required"`
	UserID string `json:"userId" binding:"required"`
	UserPW string `json:"userPw" binding:"required"`
}

//ResponseUserInfo /v1/get으로 들어온 요청에 대한 응답
type ResponseUserInfo struct {
	ID         int    `json:"_id"`
	UserID     string `json:"userId"`
	UserName   string `json:"userName"`
	UserNumber int    `json:"userNumber"`
}

//UserPost 유저 정보를 불러오는 API
//->RequestUserInfo
//<-ResponseUserInfo
//400, 401, 403, 500 에러 발생가능
func UserPost(c *gin.Context) {
	var req RequestUserInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Connect()
	defer db.Close()

	var client database.ClientInfo

	// 유효하지 않은 토큰이면 401 에러를 발생시킨다.
	if err := db.Where(
		database.ClientInfo{Token: req.Token},
	).Take(&client).Error; err != nil ||
		client.Valid != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 유효한 유저정보를 찾지못하면 403 에러를 발생시킨다.
	var user database.UserInfo
	if err := db.Where(
		database.UserInfo{UserID: req.UserID},
	).Take(&user, "UserID = ?", req.Token).Error; err != nil ||
		req.UserPW != utils.RabumsHashedPasswdToken(user.UserPW, client.Token) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// 유저정보 반환
	var res ResponseUserInfo
	res.ID = user.ID
	res.UserID = user.UserID
	res.UserName = user.UserName
	res.UserNumber = user.UserNumber
	c.JSON(http.StatusOK, res)
}
