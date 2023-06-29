package shop

import (
	"ET-order-mini-program/database/models"
	"errors"

	"gorm.io/gorm"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{db: db}
}

func (r *ShopRepository) Create(shop *models.Shop) error {
	result := r.db.Create(shop)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to create shop")
	}
	return nil
}

func (r *ShopRepository) Update(shop *models.Shop) error {
	result := r.db.Save(shop)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update goods")
	}
	return nil
}

func (r *ShopRepository) Delete(shop *models.Shop) error {
	result := r.db.Delete(shop)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete goods")
	}
	return nil
}

func (r *ShopRepository) GetById(id uint) (*models.Shop, error) {
	var shop models.Shop
	// 预加载 Goods 关系
	result := r.db.Preload("Goods").First(&shop, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("goods not found")
	}
	return &shop, nil
}

func (r *ShopRepository) GetAll() ([]models.Shop, error) {
	var shops []models.Shop
	// 预加载 Goods 关系
	result := r.db.Preload("Goods").Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}
	return shops, nil
}
