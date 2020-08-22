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
	UserNumber int32  `json:"userNumber"`
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


	// 유효하지 않은 토큰이면 401 에러를 발생시킨다.
	client, err := database.CheckClientToken(req.Token)
	if err != nil {
		switch err {
		case database.ErrRequestNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			database.RecordLog(0, 0, database.INVALIDTOKEN, database.Message{
				"token":  req.Token,
				"userId": req.UserID,
				"userPw": req.UserPW,
			})
			break
		default:
			panic(err)
		}
		return
	}

	// 유효한 유저정보를 찾지못하면 403 에러를 발생시킨다.
	user, err := database.GetUserByUserID(req.UserID)
	if err != nil {
		switch err {
		case database.ErrRequestNotFound:
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid Userinfo"})
			database.RecordLog(0, client.ID, database.USERNOTFOUND, database.Message{
				"userId": req.UserID,
				"userPw": req.UserPW,
			})
			break
		default:
			panic(err)
		}
		return
	}

	if req.UserPW != utils.RabumsHashedPasswdToken(user.UserPW, client.Token) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid Userinfo"})
		database.RecordLog(user.ID, client.ID, database.INCORRECTPASSWORD, database.Message{
			"userId": req.UserID,
			"userPw": req.UserPW,
		})
		return
	}

	// 유저정보 반환
	var res ResponseUserInfo
	res.ID = user.ID
	res.UserID = user.UserID
	res.UserName = user.UserName
	res.UserNumber = user.UserNumber
	c.JSON(http.StatusOK, res)
	database.RecordLog(user.ID, client.ID, database.LOGIN, database.Message{})
}
