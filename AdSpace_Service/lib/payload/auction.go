package payload

import "time"

type Auction struct {
	ID        int       `json:"id"`
	AdSpaceID int       `json:"adSpaceID"`
	BasePrice string    `json:"basePrice"`
	Name      string    `json:"name"`
	EndTime   time.Time `json:"endTime"`
}

type AuctionReq struct {
	AdSpaceID *int       `json:"adSpaceID"`
	EndTime   *time.Time `json:"endTime"`
}
