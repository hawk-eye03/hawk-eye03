package payload

type Bidder struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BidderReq struct {
	Name string `json:"name"`
}
