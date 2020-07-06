package services

import (
	"blog-server/models"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Dialect support for sqlite
)

var dbPool *gorm.DB

func createIfNotExist(dirPath string, filePath string) {
	if _, err = os.Stat(dirPath); os.IsNotExist(err) {
	}
}

func initDatabase() {

	databasePath, databaseDirPath := (func() (string, string) {
		currentInfo, err := user.Current()
		if err != nil {
			panic(err)
		}
		dirPath := filepath.Join(currentInfo.HomeDir, "assets")
		filePath := filepath.Join(dirPath, "database.db")
		return filePath, dirPath
	})()

	fmt.Println(databasePath)
	db, err := gorm.Open("sqlite3", databasePath)
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
