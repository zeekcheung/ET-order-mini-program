package shop

import "ET-order-mini-program/database/models"

type CreateShopDto struct {
	OwnerID      string
	Name         string
	AvatarUrl    string
	Introduction string
	Address      string
}

func (d CreateShopDto) CreateShop() *models.Shop {
	return &models.Shop{
		OwnerID:      d.OwnerID,
		Name:         d.Name,
		AvatarUrl:    d.AvatarUrl,
		Introduction: d.Introduction,
		Address:      d.Address,
	}
}
