package service

import (
	"log"

	"github.com/shreyash-tripathi/ad-service/lib/db"
	"github.com/shreyash-tripathi/ad-service/lib/db/mapper"
	"github.com/shreyash-tripathi/ad-service/lib/payload"
)

type AdSpaceService struct{}

func (AdServ *AdSpaceService) CreateAdSpace(adSpace payload.AdSpaceReq) error {
	log.Println("Inside Create Ad Space Service")
	newAdSpace := &mapper.AdSpace{
		Name:      adSpace.Name,
		BasePrice: adSpace.BasePrice,
	}

	err := db.CreateAdSpace(newAdSpace)
	if err != nil {
		return err
	}
	return nil
}
