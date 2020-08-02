package v1

//RequestUserInfo /v1/get으로 들어오는 요청 타입
type RequestUserInfo struct {
	Token  string `json:"token" binding:"required"`
	UserID string `json:"userId" binding:"required"`
	UserPW string `json:"userPw" binding:"required"`
}

//ResponseUserInfo /v1/get으로 들어온 요청에 대한 응답
type ResponseUserInfo struct {
	ID         int    `json:"_id"`
	UserID     string `json:"userId"`
	UserName   string `json:"userName"`
	UserNumber int    `json:"userNumber"`
}
