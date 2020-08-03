package v1

import (
	"net/http"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/gin-gonic/gin"
)

/*
	1. 유저가 회원가입 정보를 입력한다.
	2. /v1/user/validating 으로 유저정보를 보낸다.
	3. 인증 토큰(AES)을 만들어서 유저 이메일로 보낸다.
	4. 유저가 인증 토큰을 RequestUserCreate로 보내면 유저 생성함
*/

//RequestUserCreate /v1/user PUT 요청
type RequestUserCreate struct {
	Token string `json:"token" binding:"required"`
}

//RequestUserToken /v1/user/validating PUT 요청
type RequestUserToken struct {
	UserName string `json:"userName" binding:"required"`
	// 학생이 아닌경우 일단 -1을 받고 임의의 번호 배정
	UserNumber int    `json:"userNumber" binding:"required"`
	UserEmail  string `json:"userEmail" binding:"required"`
	UserID     string `json:"userId" binding:"required"`
	UserPW     string `json:"userPw" binding:"required"`
}

//ResponseUserPut /v1/user /v1/user/validating PUT 요청 응답
type ResponseUserPut struct {
	Ok bool `json:"ok" binding:"required"`
}

type tokneInfo struct {
	jobType string
	body    RequestUserToken
}

//UserValidatingPut /v1/user/validating PUT 요청
//->RequestUserToken
//<-ResponseUserPut
//400, 500 에러 발생 가능
func UserValidatingPut(c *gin.Context) {
	var req RequestUserToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Connect()
	defer db.Close()

	//유효하지 않은 이메일 주소이면 400 에러를 발생시킨다.
	//email := req.UserEmail

	//이미 존재하지 않는 ID이면 400 에러를 발생시킨다.

	//이미 존재하는 학번이면 400 에러를 발생시킨다.

	//토큰을 생성한다.

	//이메일을 보내는데 실패하면 500 에러를 발생시킨다.
}

//UserPut /v1/user PUT 요청
//->RequestUserCreate
//<-ResponseUserPut
//400, 409, 500에러 발생가능
func UserPut(c *gin.Context) {
	var req RequestUserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//토큰 검증에 실패하면 400 에러를 발생시킨다.

	//DB에 이미 존재하는 데이터면 409 에러를 발생시킨다.
}
