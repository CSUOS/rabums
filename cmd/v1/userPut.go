package v1

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/CSUOS/rabums/cmd/database"
	"github.com/CSUOS/rabums/cmd/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	db := database.Connect()
	defer db.Close()

	//유효하지 않은 이메일 주소이면 400 에러를 발생시킨다.
	var user database.UserInfo
	email := req.UserEmail
	if strings.HasSuffix(email, "@uos.ac.kr") == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않는 이메일주소"})
		return
	}

	//이미 존재하는 이메일 주소이면 400 에러를 발생시킨다.
	if err := db.Where(
		database.UserInfo{UserEmail: email, Available: true},
	).Take(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "이미 가입된 이메일주소"})
		return
	}

	//이미 존재하는 ID이면 400 에러를 발생시킨다.
	if err := db.Where(
		database.UserInfo{UserID: req.UserID},
	).Take(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "이미 사용중인 ID"})
		return
	}

	//학생이 아니면 적당한 번호를 배정한다.
	for req.UserNumber == -1 {
		req.UserNumber = (rand.Int31() % 1000000000) + 1000000000
		var count int
		if err := db.Where(
			database.UserInfo{UserNumber: req.UserNumber},
		).Count(&count).Error; err != nil {
			panic(err)
		}
		if count != 0 {
			req.UserNumber = -1
		}
	}

	//이미 존재하는 학번이면 400 에러를 발생시킨다.
	if err := db.Where(database.UserInfo{
		UserNumber: req.UserNumber, Available: true,
	}).Take(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "이미 가입된 학번"})
		return
	}

	//토큰을 생성한다.
	db.FirstOrCreate(&user, database.UserInfo{
		UserID: req.UserID,
	})
	user.UserID = req.UserID
	user.UserPW = req.UserPW
	user.UserEmail = req.UserEmail
	user.UserName = req.UserName
	user.UserNumber = req.UserNumber
	user.Available = false
	if err := db.Save(&user).Error; err != nil {
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

	db := database.Connect()
	defer db.Close()

	user := database.UserInfo{}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where(database.UserInfo{
		UserID:     fmt.Sprintf("%v", claim["userId"]),
		UserNumber: int32(claim["userNumber"].(float64)),
		UserEmail:  fmt.Sprintf("%v", claim["userEmail"]),
	}).Not(
		database.UserInfo{Available: true},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			//DB에 이미 존재하는 데이터면 409 에러를 발생시킨다.
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Available = true

	if err := db.Save(&user).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, ResponseUserPut{Ok: true})
}
