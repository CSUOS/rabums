package database

import (
	"github.com/jinzhu/gorm"
	// mysql 플러그인
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Connect DB랑 연결함
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "rabums:mhmvAfBPWcF4NX1y@(doctor.iptime.org)/rabums?charset=utf8mb4&parseTime=True")
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
