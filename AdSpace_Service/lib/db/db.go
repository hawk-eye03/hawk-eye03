package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shreyash-tripathi/ad-service/lib/db/mapper"
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

func CreateAuction(auction *mapper.Auction) error {
	log.Println("Inside Create Auction DB call")
	dbErr := DB.Create(auction).Error
	if dbErr != nil {
		return dbErr
	}
	return nil
}

func CreateAdSpace(adSpace *mapper.AdSpace) error {
	log.Println("Inside Create Ad Space DB call")
	dbErr := DB.Create(adSpace).Error
	if dbErr != nil {
		return dbErr
	}
	return nil
}

func GetListOfAuctions() ([]mapper.AuctionAdSpaceDetails, error) {
	log.Println("Inside Get List of Auctions DB call")

	var auctionAdSpaceDetailsList []mapper.AuctionAdSpaceDetails
	auction := mapper.Auction{}

	dbErr := DB.Model(&auction).Select("auctions.id, auctions.ad_space_ref_id, auctions.end_time, ad_spaces.base_price, ad_spaces.name").Joins("left join ad_spaces on ad_spaces.id = auctions.ad_space_ref_id").Scan(&auctionAdSpaceDetailsList).Error
	if dbErr != nil {
		return nil, dbErr
	}

	return auctionAdSpaceDetailsList, nil
}
