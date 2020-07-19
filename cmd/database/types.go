package types

import "time"

/*
[TABLE] ClientInfo

[TABLE] UserInfo

[TABLE] Log
*/

//ClientInfo 클라이언트정보 DB schema
type ClientInfo struct {
	ID    int
	Name  string
	Token string
	Valid bool
}

//UserType 사용자 타입
type UserType int8

//사용자 종류
const (
	ADMINISTRATOR UserType = -1
	STUDENT       UserType = iota
	PROFESSOR
	INSTRUCTOR
	EXTERNAL
)

//UserInfo 사용자정보 DB schema
type UserInfo struct {
	ID        int
	Name      string
	StudNum   int
	Email     string
	UserID    string
	UserPW    string
	CreatedAt time.Time
}

//EventType 발생가능한 이벤트 타입
type EventType int

//사용자 종류
const (
	UNKOWN    EventType = -1
	CREATED   EventType = 100
	MODIFIED  EventType = 101
	DELETED   EventType = 102
	LOGIN     EventType = 200
	LOGOUT    EventType = 201
	LOGINFAIL EventType = 202
	QUERYLOG  EventType = 300
)

//Log 이벤트발생시 기록남기는 DB schema
type Log struct {
	User      int
	Client    int
	Event     EventType
	Message   string
	CreatedAt time.Time
}
