package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/CSUOS/rabums/cmd/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
	1. 유저가 회원가입 정보를 입력한다.
	2. /v1/user/validating 으로 유저정보를 보낸다.
	3. 인증 토큰(JWT)을 만들어서 유저 이메일로 보낸다.
	4. 유저가 인증 토큰을 RequestUserCreate로 보내면 유저 생성함
*/

//RequestUserCreate /v1/user/request/register PUT 요청
type RequestUserCreate struct {
	Token string `json:"token" binding:"required"`
}

//RequestUserToken /v1/user/request/token PUT 요청
type RequestUserToken struct {
	UserName string `json:"userName" binding:"required"`
	// 학생이 아닌경우 일단 -1을 받고 임의의 번호 배정
	UserNumber int32  `json:"userNumber" binding:"required"`
	UserEmail  string `json:"userEmail" binding:"required"`
	UserID     string `json:"userId" binding:"required"`
	UserPW     string `json:"userPw" binding:"required"`
}

//ResponseUserPut PUT 요청 응답
type ResponseUserPut struct {
	Ok bool `json:"ok" binding:"required"`
}

//UserReqTokenPut /v1/user/request/token PUT 요청
//->RequestUserToken
//<-ResponseUserPut
//400, 500 에러 발생 가능
func UserReqTokenPut(c *gin.Context) {
	var req RequestUserToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//유효하지 않은 이메일 주소이면 400 에러를 발생시킨다.
	email := req.UserEmail
	if strings.HasSuffix(email, "@uos.ac.kr") == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않는 이메일주소"})
		return
	}

	//이미 존재하는 이메일 주소이면 400 에러를 발생시킨다.
	user, err := database.GetUserByEmail(email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "이미 가입된 이메일주소"})
		return
	}
	//비활성 상태의 정보는 제거
	database.DeleteUserByUserEmail(req.UserEmail)

	//이미 존재하는 ID이면 400 에러를 발생시킨다.
	user, err = database.GetUserByUserIDIgnoreAvailability(req.UserID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "이미 사용중인 ID"})
		return
	}

	//학생이 아니면 적당한 번호를 배정한다.
	if req.UserNumber == -1 {
		req.UserNumber = 0
	} else {
		//이미 존재하는 학번이면 400 에러를 발생시킨다.
		user, err = database.GetUserByUserNumber(req.UserNumber)
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "이미 가입된 학번"})
			return
		}
	}

	//토큰을 생성한다.
	user = &database.UserInfo{}

	user.UserID = req.UserID
	user.UserPW = req.UserPW
	user.UserEmail = req.UserEmail
	user.UserName = req.UserName
	user.UserNumber = req.UserNumber
	user.Available = false

	if err := database.UpdateUser(user); err != nil {
		panic(err)
	}

	token := utils.Obeject2JWT(jwt.MapClaims{
		"userId":     user.UserID,
		"userNumber": user.UserNumber,
		"userEmail":  user.UserEmail,
	})
	//이메일을 보내는데 실패하면 500 에러를 발생시킨다.
	if err := utils.SendTokenViaEmail(user.UserEmail, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "이메일 보내기에 실패하였습니다. 관리자에게 문의해주세요. :-("})
		return
	}

	c.JSON(http.StatusOK, ResponseUserPut{
		Ok: true,
	})
	return
}

//UserReqRegisterPut /v1/user/request/register PUT 요청
//->RequestUserCreate
//<-ResponseUserPut
//400, 409, 500에러 발생가능
func UserReqRegisterPut(c *gin.Context) {
	var req RequestUserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//토큰 검증에 실패하면 400 에러를 발생시킨다.
	claim, err := utils.JWT2Object(req.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.GetUserByEmailNotAvailable(fmt.Sprintf("%v", claim["userEmail"]))

	if err != nil {
		switch err {
		case database.ErrRequestNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		default:
			panic(err)
		}
	}

	if user.UserNumber != int32(claim["userNumber"].(float64)) || user.UserID != fmt.Sprintf("%v", claim["userId"]) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.UserNumber == 0 {
		user.UserNumber = int32(user.ID) + 1000000000
	}

	user.Available = true

	err = database.UpdateUser(user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, ResponseUserPut{Ok: true})
}
