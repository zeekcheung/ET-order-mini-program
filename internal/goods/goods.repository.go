package goods

import (
	"ET-order-mini-program/database/models"
	"errors"

	"gorm.io/gorm"
)

type GoodsRepository struct {
	db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepository {
	return &GoodsRepository{db: db}
}

func (r *GoodsRepository) Create(goods *models.Goods) error {
	result := r.db.Create(goods)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to create goods")
	}
	return nil
}

func (r *GoodsRepository) Update(goods *models.Goods) error {
	result := r.db.Save(goods)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update goods")
	}
	return nil
}

func (r *GoodsRepository) Delete(goods *models.Goods) error {
	result := r.db.Delete(goods)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete goods")
	}
	return nil
}

func (r *GoodsRepository) GetById(id string) (*models.Goods, error) {
	var goods models.Goods
	// 预加载 GoodsSpecs 关系
	result := r.db.Preload("GoodsSpecs").First(&goods, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("goods not found")
	}
	return &goods, nil
}

func (r *GoodsRepository) GetAll() ([]models.Goods, error) {
	var goods []models.Goods
	// 预加载 GoodsSpecs 关系
	result := r.db.Preload("GoodsSpecs").Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}
	return goods, nil
}
