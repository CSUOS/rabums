package server

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

//GetReason return string format to Reason struct
func GetReason(format string, a ...interface{}) Reason {
	return Reason{Reason: fmt.Sprintf(format, a...)}
}

func send(w http.ResponseWriter, statusCode int, data interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}

//GenerateNewToken 임의의 토큰을 생성해서 반환한다.
func GenerateNewToken() string {
	t := fmt.Sprintf("%d", rand.Int31())
	base64 := base64.StdEncoding.EncodeToString([]byte(t))
	h := sha256.New()
	h.Write([]byte(base64))
	md := h.Sum(nil)
	return hex.EncodeToString(md)
}
