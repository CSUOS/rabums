package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
	DeletedAt  time.Time
}

//Create user 생성
func (u *UserInfo) Create() error {
	if u == nil {
		return ErrRequestCanNotBeNil
	}
	db := GetDB()
	if err := db.Create(u).Error; err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//Get userId랑 일치하는거 데이터베이스에서 가져와줌
func (u *UserInfo) Get(UserID string) error {
	db := GetDB()
	if err := db.Where(UserInfo{UserID: UserID}).Take(u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrRequestNotFound
		} else {
			return ErrDatabaseUnavailable
		}

	}
	return nil
}

//Fetch 현재 저장된 정보와 일치하는정보를 DB에서 가져옴
func (u *UserInfo) Fetch() error {
	db := GetDB()
	if err := db.Where(u).Take(u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrRequestNotFound
		} else {
			return ErrDatabaseUnavailable
		}

	}
	return nil
}

//Update 수정된 사항 데이터베이스에 저장해줌
func (u *UserInfo) Update() error {
	db := GetDB()
	if err := db.Save(u).Error; err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//GetLogs 현재 유저의 로그정보를 불러온다.
func (u *UserInfo) GetLogs(Offset int, Limit int) (*[]Log, error) {
	output := new([]Log)
	db := GetDB()
	if err := db.Where(Log{User: u.ID}).Offset(Offset).Limit(Limit).Order("created_at desc").Find(output).Error; err != nil {
		return nil, ErrDatabaseUnavailable
	}
	return output, nil
}

//Delete 삭제
func (u *UserInfo) Delete() error {
	db := GetDB()
	tmp := &UserInfo{}
	if err := db.Where(u).Delete(tmp).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrRequestNotFound
		} else {
			return ErrDatabaseUnavailable
		}
	}
	return nil
}
