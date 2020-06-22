package db

import (
	"fmt"

	"github.com/SolidShake/GoMonitoring/internal/config"
	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDB() {
	cnf := config.GetConfig()
	dbConfig := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Host,
		cnf.Database.DbName,
	)

	d, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		panic(err)
	}

	db = d
}

func CloseDB() error {
	return db.Close()
}

func GetDB() *gorm.DB {
	return db
}
