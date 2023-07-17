package payload

type AdSpace struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BasePrice string `json:"basePrice"`
}

type AdSpaceReq struct {
	Name      string `json:"name"`
	BasePrice string `json:"basePrice"`
}
