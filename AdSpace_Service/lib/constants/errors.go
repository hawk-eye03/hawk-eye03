package constants

const (
	ErrorAdSpaceIDEmpty    = 1
	ErrorAdSpaceIDEmptyMsg = "Ad Space ID cannot be empty"

	ErrorAuctionEndTimeEmpty    = 2
	ErrorAuctionEndTimeEmptyMsg = "Auction end time cannot be empty"

	ErrorInvalidAuctionRequestBody    = 3
	ErrorInvalidAuctionRequestBodyMsg = "Auction request body is invalid"

	ErrorFromDBInCreatingNewAuction    = 4
	ErrorFromDBInCreatingNewAuctionMsg = "DB Error while creating new auction: %v"

	ErrorInvalidAdSpaceRequestBody    = 5
	ErrorInvalidAdSpaceRequestBodyMsg = "Ad Space request body is invalid"

	ErrorAdSpaceNameEmpty    = 6
	ErrorAdSpaceNameEmptyMsg = "Ad Space name cannot be empty"

	ErrorAdSpaceBasePriceEmpty    = 7
	ErrorAdSpaceBasePriceEmptyMsg = "Auction ad space base price cannot be empty"

	ErrorFromDBInCreatingNewAdSpace    = 8
	ErrorFromDBInCreatingNewAdSpaceMsg = "DB Error while creating new ad space: %v"

	ErrorFromDBInFetchingAllAuctions    = 9
	ErrorFromDBInFetchingAllAuctionsMsg = "DB Error while fetching all auctions: %v"
)
