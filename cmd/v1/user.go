package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LEE-WAN/RABUMS/cmd/util"

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

//UserPost 유저 정보를 불러오는 API
//userId, userPw, token 으로 사용자 정보를 받아옴.
//_id, userId, userName, userNumber 를 반환함.
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
		req.UserPW != util.RabumsHashedPasswdToken(user.UserPW, client.Token) {
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
