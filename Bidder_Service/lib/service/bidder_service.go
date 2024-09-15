package service

import (
	"github.com/shreyash-tripathi/bidder-service/lib/db"
	"github.com/shreyash-tripathi/bidder-service/lib/db/mapper"
	"github.com/shreyash-tripathi/bidder-service/lib/payload"
)

type BidderService struct{}

func (bs *BidderService) CreateBidder(bidder payload.BidderReq) error {
	newBidder := &mapper.Bidder{
		Name: bidder.Name,
	}
	err := db.CreateBidder(newBidder)
	if err != nil {
		return err
	}
	return err
}
