package types

//RequestUserInfo /v1/get으로 들어오는 요청 타입
type RequestUserInfo struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
	UserPW string `json:"userPw"`
}

//ResponseUserInfo /v1/get으로 들어온 요청에 대한 응답
type ResponseUserInfo struct {
	ID      int    `json:"_id"`
	UserID  string `json:"userId"`
	Name    string `json:"name"`
	StudNum int    `json:"studNum"`
}

