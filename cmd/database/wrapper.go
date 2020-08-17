package database

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

//Message Logging 용 Message
type Message map[string]interface{}

//RecordLog 기록용
func RecordLog(user int, client int, event EventType, msg Message) error {
	db := GetDB()
	message, _ := json.Marshal(msg)
	db.Create(&Log{
		User:    user,
		Client:  client,
		Event:   event,
		Message: string(message),
	})

	return nil
}

//GetClientInfo client 정보 가져옴
func GetClientInfo(clientID string) (*ClientInfo, error) {
	result := &ClientInfo{}
	db := GetDB()
	if err := db.Where(ClientInfo{ClientID: clientID}).
		FirstOrCreate(result).Error; err != nil {
		return nil, ErrDatabaseUnavailable
	}
	return result, nil
}

//UpdateClientInfo client 정보 갱신
func UpdateClientInfo(clientInfo *ClientInfo) error {
	db := GetDB()
	if err := db.Save(clientInfo).Error; err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//CheckClientToken 유효한 토큰인지 확인
func CheckClientToken(tkn string) (*ClientInfo, error) {
	db := GetDB()
	client := &ClientInfo{}
	if err := db.Where(
		ClientInfo{
			Token: tkn,
			Valid: true,
		},
	).Take(&client).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return client, nil
}

//GetUser ID로 유저 정보를 불러온다.
func GetUser(ID int) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	if err := db.Where(
		UserInfo{ID: ID},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return user, nil
}

//GetUserByUserID userId로 유저 정보를 불러온다.
func GetUserByUserID(userID string) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	if err := db.Where(
		UserInfo{
			UserID:    userID,
			Available: true,
		},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return user, nil
}

//GetUserByUserIDIgnoreAvailability GetUserByUserID와 같으나, 비활성 상태인 계정도 확인
func GetUserByUserIDIgnoreAvailability(userID string) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	if err := db.Where(
		UserInfo{
			UserID: userID,
		},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return user, nil
}

//GetUserByEmail userEmail로 유저 정보를 불러온다.
func GetUserByEmail(userEmail string) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	if err := db.Where(
		UserInfo{
			UserEmail: userEmail,
			Available: true,
		},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return user, nil
}

//GetUserByEmailNotAvailable userEmail로 유저 정보를 불러온다.
func GetUserByEmailNotAvailable(userEmail string) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	if err := db.Where(
		UserInfo{
			UserEmail: userEmail,
			Available: false,
		},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return user, nil
}

//GetUserByUserNumber userNumber로 유저 정보를 불러온다.
func GetUserByUserNumber(userNumber int32) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	if err := db.Where(
		UserInfo{UserNumber: userNumber},
	).Take(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRequestNotFound
		}
		return nil, ErrDatabaseUnavailable
	}
	return user, nil
}

//CreateUser 새로운 유저 생성
func CreateUser(userInfo *UserInfo) error {
	db := GetDB()
	if err := db.Create(userInfo); err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//UpdateUser 유저 정보 갱신
func UpdateUser(userInfo *UserInfo) error {
	db := GetDB()
	if err := db.Save(userInfo).Error; err != nil {
		return ErrDatabaseUnavailable
	}
	return nil
}

//DeleteUserByUserEmail 특정 이메일의 유저 삭제
func DeleteUserByUserEmail(userEmail string) error {
	db := GetDB()
	tmp := &UserInfo{}
	if err := db.Where(UserInfo{UserEmail: userEmail}).Delete(tmp).Error; err != nil {
		return ErrRequestNotFound
	}
	return nil
}
