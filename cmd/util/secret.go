package util

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var masterKey string

//GenerateSecret 현 세션에서 사용할 마스터키를 생성한다.
func GenerateSecret() {
	t := time.Now().String()
	masterKey = RabumsHASH(t)
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
