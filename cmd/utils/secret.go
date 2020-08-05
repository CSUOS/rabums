package utils

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var masterKey string
var secretKey string

//GenerateSecret 현 세션에서 사용할 마스터키를 생성한다.
func GenerateSecret() {
	t := time.Now().String()
	secretKey = RabumsHASH(t)
	masterKey = RabumsHASH(secretKey)
	log.Printf("New masterkey is %q", masterKey)
}

//GenerateNewToken 임의의 토큰을 생성해서 반환한다.
func GenerateNewToken() string {
	t := fmt.Sprintf("%d", rand.Intn(1000000))
	return RabumsHASH(t)
}

//CheckIsMasterkey 들어온 입력이 마스터키인지 확인한다.
func CheckIsMasterkey(str string) bool {
	return str == masterKey
}

//Obeject2JWT JWT 가져오는 함수
func Obeject2JWT(o jwt.MapClaims) string {
	o["expiers_at"] = time.Now().Add(time.Hour).Format(time.RFC3339)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, o)
	tokenString, _ := token.SignedString([]byte(secretKey))
	return tokenString
}

//JWT2Object JWT 검증하는 함수
func JWT2Object(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		t, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", claims["expiers_at"]))
		if err != nil {
			return nil, err
		}
		if time.Now().Before(t) {
			return claims, nil
		}
		return nil, fmt.Errorf("Out of the date")
	}
	return nil, err
}
