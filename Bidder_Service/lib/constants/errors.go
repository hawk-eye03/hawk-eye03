package constants

const (
	ErrorInvalidBidderRequestBody    = 1
	ErrorInvalidBidderRequestBodyMsg = "Bidder request body is invalid"

	ErrorBidderNameEmpty    = 2
	ErrorBidderNameEmptyMsg = "Biddder name cannot be empty"

	ErrorFromDBInCreatingNewBidder    = 3
	ErrorFromDBInCreatingNewBidderMsg = "DB Error while creating new bidder: %v"

	ErrorFromDBInFetchingAllAuctions    = 4
	ErrorFromDBInFetchingAllAuctionsMsg = "DB Error while fetching all auctions: %v"
)
