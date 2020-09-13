package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

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

//Create 새로운 ClientInfo 생성
func (c *ClientInfo) Create() error {
	if c == nil {
		return ErrRequestCanNotBeNil
	}
	db := GetDB()
	if err := db.Create(c).Error; err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//Get clientId랑 일치하는거 데이터베이스에서 가져와줌
func (c *ClientInfo) Get(clientID string) error {
	db := GetDB()
	if err := db.Where(ClientInfo{ClientID: clientID}).Take(c).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrRequestNotFound
		} else {
			return ErrDatabaseUnavailable
		}

	}
	return nil
}

//Update 수정된 사항 데이터베이스에 저장해줌
func (c *ClientInfo) Update() error {
	db := GetDB()
	if err := db.Save(c).Error; err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//Delete 삭제
func (c *ClientInfo) Delete() error {
	db := GetDB()
	tmp := &UserInfo{}
	if err := db.Where(c).Delete(tmp).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrRequestNotFound
		} else {
			return ErrDatabaseUnavailable
		}
	}
	return nil
}

//GetLogs 현재 유저의 로그정보를 불러온다.
func (c *ClientInfo) GetLogs(Offset int, Limit int) (*[]Log, error) {
	output := new([]Log)
	db := GetDB()
	if err := db.Where(Log{Client: c.ID}).Offset(Offset).Limit(Limit).Order("created_at desc").Find(output).Error; err != nil {
		return nil, ErrDatabaseUnavailable
	}
	return output, nil
}
