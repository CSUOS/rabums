package database

import "time"

/*
[TABLE] ClientInfo

[TABLE] UserInfo

[TABLE] Log
*/



//Log 이벤트발생시 기록남기는 DB schema
type Log struct {
	ID        int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	User      int       `gorm:"index:user"`
	Client    int       `gorm:"index:client"`
	Event     EventType `gorm:"index:event"`
	Message   string
	CreatedAt time.Time
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

//EventType 발생가능한 이벤트 타입
type EventType int

//사용자 종류
const (
	UNKOWN   EventType = -1
	ACCESSED EventType = 100
	CREATED  EventType = 101
	UPDATED  EventType = 102
	DELETED  EventType = 103
	LOGIN    EventType = 200
	LOGOUT   EventType = 201

	INVALIDCLIENTPW   EventType = 401
	INVALIDTOKEN      EventType = 402
	USERNOTFOUND      EventType = 403
	INCORRECTPASSWORD EventType = 404
)
