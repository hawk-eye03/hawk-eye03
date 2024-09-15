package mapper

import "time"

type Auction struct {
	ID           int
	AdSpaceRefID int
	EndTime      time.Time
}

type AuctionAdSpaceDetails struct {
	ID           int
	AdSpaceRefID int
	BasePrice    string
	Name         string
	EndTime      time.Time
}
