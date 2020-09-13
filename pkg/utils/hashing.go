package utils

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
)

//RabumsHASH 라붐스에서 자체적으로 사용하는 해싱 함수
func RabumsHASH(data string) string {
	base64 := b64.StdEncoding.EncodeToString([]byte(data))
	h := sha256.New()
	h.Write([]byte(base64))
	md := h.Sum(nil)
	return hex.EncodeToString(md)
}

//RabumsHashedPasswdToken 해싱된 패스워드랑 토큰 원본 조합으로 토큰에 알맞은 해시값 구해냄
func RabumsHashedPasswdToken(hashedpassword string, token string) string {
	return RabumsHASH(RabumsHASH(token) + hashedpassword)
}

//RabumsTokenHash 주어진 토큰에 맞는 해시값을 구해내는 함수
//샘플용이고 실질적으로 사용 안함
func RabumsTokenHash(password string, token string) string {
	hashedPassword := RabumsHASH(RabumsHASH(password))
	hashedToken := RabumsHASH(token)
	return RabumsHASH(hashedToken + hashedPassword)
}
