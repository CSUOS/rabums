package v1

import "github.com/gin-gonic/gin"

//Get 사용자 정보를 불러오는 API
//userId, userPw, token 으로 사용자 정보를 받아옴.
//_id, userId, name, studNum 받을 수 있음
//400, 401, 403, 500 에러 발생가능
func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"HELLO": "WORLD",
	})
}
