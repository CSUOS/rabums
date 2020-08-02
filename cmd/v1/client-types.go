package v1

//RequestClientUpdate /v1/client PUT으로 들어오는 요청
type RequestClientUpdate struct {
	MasterKey string `json:"masterkey" binding:"required"`
	ClientID  string `json:"clientId" binding:"required"`
	ClientPW  string `json:"clientPw" binding:"required"`
	Link      string `json:"link" binding:"required"`
	Valid     bool   `json:"valid"`
}

//ResponseClientUpdate /v1/client PUT 요청에 대한 응답
type ResponseClientUpdate struct {
	Token string `json:"token"`
}
