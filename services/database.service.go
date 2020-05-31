package services

import (
	"blog-server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Dialect support for sqlite
)

var dbPool *gorm.DB

func initDatabase() {
	db, err := gorm.Open("sqlite3", ":memory:")
	db.DB().SetMaxIdleConns(10)

	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.User{})

	dbPool = db
	if err != nil {
		panic(err)
	}
}

/*
GetDB function is to get a connection from pool
*/
func GetDB() *gorm.DB {
	if dbPool == nil {
		initDatabase()
	}
	return dbPool
}
