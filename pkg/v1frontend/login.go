package v1frontend

import (
	"fmt"
	"net/http"

	"github.com/CSUOS/rabums/pkg/database"
	"github.com/CSUOS/rabums/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type requestLogin struct {
	UserID string `json:"userId" binding:"required"`
	UserPW string `json:"userPw" binding:"required"`
}

type responseLogin struct {
	UserID     string `json:"userId"`
	UserName   string `json:"userName"`
	UserNumber int32  `json:"userNumber"`
	UserEmail  string `json:"userEmail"`
}

//GetToken Token 가져옴
func GetToken(c *gin.Context) {
	rabums := database.ClientInfo{}
	if err := rabums.Get("rabums"); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"token": utils.RabumsHASH(rabums.Token)})
}

//LogoutHandler 로그아웃
func LogoutHandler(c *gin.Context) {
	c.SetCookie("auth", "", 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

//LoginHandler POST Login
func LoginHandler(c *gin.Context) {
	var req requestLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := database.UserInfo{}
	rabums := database.ClientInfo{}
	if err := rabums.Get("rabums"); err != nil {
		panic(err)
	}

	if err := user.Get(req.UserID); err != nil {
		switch err {
		case database.ErrRequestNotFound:
			c.JSON(http.StatusForbidden, gin.H{"Error": "incorrect id or pw"})
			database.RecordLog(0, rabums.ID, database.USERNOTFOUND, database.Message{
				"userId": req.UserID,
				"userPw": req.UserPW,
			})
			return
		case database.ErrDatabaseUnavailable:
			panic(err)
		}
	}

	if req.UserPW != utils.RabumsHashedPasswdToken(user.UserPW, rabums.Token) {
		c.JSON(http.StatusForbidden, gin.H{"Error": "incorrect id or pw"})
		database.RecordLog(user.ID, rabums.ID, database.INCORRECTPASSWORD, database.Message{
			"userId": req.UserID,
			"userPw": req.UserPW,
		})
		return
	}
	database.RecordLog(user.ID, rabums.ID, database.LOGIN, nil)
	s := utils.Obeject2JWT(jwt.MapClaims{
		"userId": user.UserID,
	})
	c.SetCookie("auth", s, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, responseLogin{
		UserName: user.UserName,
		UserID: user.UserID,
		UserEmail: user.UserEmail,
		UserNumber: user.UserNumber,
	})
}

//AuthMiddleware 인증되지 않은 요청은 401을 리턴
func AuthMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("auth")
	if err != nil {
		c.SetCookie("auth", "", 3600, "/", "", false, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
		c.Abort()
		return
	}
	claim, err := utils.JWT2Object(cookie.Value)
	if err != nil || claim.Valid() != nil {
		cookie.Value = ""
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	c.Set("userId", claim["userId"])

	s := utils.Obeject2JWT(jwt.MapClaims{
		"userId": claim["userId"],
	})
	c.SetCookie("auth", s, 3600, "/", "", false, false)
}

//getUserID context에서 userId 가져옴
func getUserID(c *gin.Context) (string, error) {
	tmp := fmt.Sprintf("%s", c.MustGet("userId"))
	if tmp == "" {
		return "", fmt.Errorf("no user in context")
	}
	return tmp, nil
}
