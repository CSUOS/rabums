package database

import (
	"github.com/CSUOS/rabums/cmd/utils"
	"github.com/jinzhu/gorm"

	// mysql 플러그인
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Connect DB랑 연결함
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", utils.Config.MYSQLUri)
	if err != nil {
		panic(err)
	}
	return db
	// defer db.Close()
}

//DBInit 데이터베이스 초기화
func DBInit() {
	db := Connect()
	defer db.Close()

	db.AutoMigrate(&ClientInfo{})
	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&Log{})
}
