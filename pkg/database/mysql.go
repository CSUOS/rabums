package database

import (
	"time"

	"github.com/CSUOS/rabums/pkg/utils"
	"github.com/jinzhu/gorm"

	// mysql 플러그인
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//Ping 접속 가능여부 확인용
func Ping() error {
	db, err := gorm.Open("mysql", utils.Config.MYSQLUri)
	db.Close()
	return err
}

//GetDB DB인터페이스 소환!
func GetDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open("mysql", utils.Config.MYSQLUri)
		if err != nil {
			panic(err)
		}
		db.DB().SetConnMaxLifetime(30 * time.Minute)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
	}
	return db
}

//DBInit 데이터베이스 초기화
func DBInit() {
	GetDB()

	db.AutoMigrate(&ClientInfo{})
	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&Log{})

	c := ClientInfo{}
	if err := c.Get("rabums"); err != nil {
		c = ClientInfo{
			ClientID:    "rabums",
			ClientPW:    "nil",
			Link:        "https://rabums.csuos.ml",
			Description: "Rabums에서 자동으로 생성한 기본 계정입니다.",
			Valid:       true,
			Token:       utils.GenerateNewToken(),
		}
		err = c.Create()
		if err != nil {
			panic("DB Problem")
		}
	}
}

//Close DB 접속 끊기
func Close() {
	db.Close()
}
