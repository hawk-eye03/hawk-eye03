package service

import (
	"log"

	"github.com/shreyash-tripathi/ad-service/lib/db"
	"github.com/shreyash-tripathi/ad-service/lib/db/mapper"
	"github.com/shreyash-tripathi/ad-service/lib/payload"
)

type AuctionService struct{}

func (auctionServ *AuctionService) CreateAuction(auction payload.AuctionReq) error {
	log.Println("Inside Create Auction Service")
	newAuction := &mapper.Auction{
		AdSpaceRefID: *auction.AdSpaceID,
		EndTime:      *auction.EndTime,
	}

	err := db.CreateAuction(newAuction)
	if err != nil {
		return err
	}
	return nil
}

func (auctionServ *AuctionService) GetAllAuctions() ([]payload.Auction, error) {
	log.Println("Inside Get All Auctions Service")

	auctions, err := db.GetListOfAuctions()
	if err != nil {
		return nil, err
	}

	return prepareAuctionsPayload(auctions), nil
}

func prepareAuctionsPayload(auctions []mapper.AuctionAdSpaceDetails) []payload.Auction {
	var auctionsResp []payload.Auction
	for _, auction := range auctions {
		auctionData := payload.Auction{
			ID:        auction.ID,
			AdSpaceID: auction.AdSpaceRefID,
			Name:      auction.Name,
			BasePrice: auction.BasePrice,
			EndTime:   auction.EndTime,
		}
		auctionsResp = append(auctionsResp, auctionData)
	}
	return auctionsResp
}
