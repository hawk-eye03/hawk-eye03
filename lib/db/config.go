package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() error {
	d, err := gorm.Open("mysql", "root:Welcome@2020@/Auction?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

	err = db.DB().Ping()
	if err != nil {
		log.Println("Failed to ping MySQL database. Error :", err)
		return err
	}
	log.Println("Successfully connected to MySQL database!")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
