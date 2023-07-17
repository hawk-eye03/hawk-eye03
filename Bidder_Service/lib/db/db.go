package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shreyash-tripathi/bidder-service/lib/db/mapper"
)

var (
	DB *gorm.DB
)

func Connect() error {
	db, err := gorm.Open("mysql", "root:Welcome@2020@/Auction_Service?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = db
	err = DB.DB().Ping()
	if err != nil {
		log.Println("Failed to ping MySQL database. Error :", err)
		return err
	}
	log.Println("Successfully connected to MySQL database!")
	return nil
}

func CreateBidder(bidder *mapper.Bidder) error {
	log.Println("Inside Create Bidder DB call")
	dbErr := DB.Create(bidder).Error
	if dbErr != nil {
		return dbErr
	}
	return nil
}
