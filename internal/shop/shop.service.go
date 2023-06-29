package shop

import (
	"ET-order-mini-program/database/models"

	"gorm.io/gorm"
)

type ShopService struct {
	shopRepo *ShopRepository
}

func NewShopService(db *gorm.DB) *ShopService {
	return &ShopService{
		shopRepo: NewShopRepository(db),
	}
}

func (s *ShopService) CreateShop(dto *CreateShopDto) (*models.Shop, error) {
	// FIXME: 店铺是否需要考虑重名等问题？
	shop := dto.CreateShop()
	shop.Grade = 0
	return shop, s.shopRepo.Create(shop)
}

func (s *ShopService) GetShopById(id uint) (*models.Shop, error) {
	return s.shopRepo.GetById(id)
}

func (s *ShopService) GetAllShops() ([]models.Shop, error) {
	return s.shopRepo.GetAll()
}

func (s *ShopService) UpdateShop(shop *models.Shop) error {
	return s.shopRepo.Update(shop)
}

func (s *ShopService) DeleteShop(id uint) error {
	shop, _ := s.GetShopById(id)
	return s.shopRepo.Delete(shop)
}
