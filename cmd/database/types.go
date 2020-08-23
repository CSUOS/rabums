package database

import "time"

/*
[TABLE] ClientInfo

[TABLE] UserInfo

[TABLE] Log
*/

//ClientInfo 클라이언트정보 DB schema
type ClientInfo struct {
	ID          int       `gorm:"PRIMARY_KEY;unique_index;AUTO_INCREMENT" json:"id"`
	ClientID    string    `gorm:"PRIMARY_KEY;unique_index" json:"clientId"`
	ClientPW    string    `gorm:"size:255" json:"clientPw"`
	Link        string    `gorm:"size:255" json:"link"`
	Token       string    `gorm:"size:255" json:"token"`
	Description string    `gorm:"size:4096" json:"description"`
	Valid       bool      `json:"valid"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
	UNKOWN   EventType = -1
	ACCESSED EventType = 100
	CREATED  EventType = 101
	UPDATED  EventType = 102
	DELETED  EventType = 103
	LOGIN    EventType = 200
	LOGOUT   EventType = 201

	INVALIDMASTERKEY  EventType = 401
	INVALIDTOKEN      EventType = 402
	USERNOTFOUND      EventType = 403
	INCORRECTPASSWORD EventType = 404
)
