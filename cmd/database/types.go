package database

import "time"

/*
[TABLE] ClientInfo

[TABLE] UserInfo

[TABLE] Log
*/

//ClientInfo 클라이언트정보 DB schema
type ClientInfo struct {
	ID        int    `gorm:"PRIMARY_KEY;unique_index;AUTO_INCREMENT"`
	ClientID  string `gorm:"PRIMARY_KEY;unique_index"`
	ClientPW  string `gorm:"size:255"`
	Link      string `gorm:"size:255"`
	Token     string `gorm:"size:255"`
	Valid     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//UserInfo 사용자정보 DB schema
type UserInfo struct {
	ID         int    `gorm:"PRIMARY_KEY;unique_index;AUTO_INCREMENT"`
	UserName   string `gorm:"index:userName"`
	UserNumber int32  `gorm:"index:userNumber"`
	UserEmail  string
	UserID     string `gorm:"unique_index"`
	UserPW     string
	Available  bool `gorm:"index:availableUser"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

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
	UNKOWN    EventType = -1
	CREATED   EventType = 100
	UPDATED   EventType = 101
	DELETED   EventType = 102
	LOGIN     EventType = 200
	LOGOUT    EventType = 201
	LOGINFAIL EventType = 202
	QUERYLOG  EventType = 300
)
