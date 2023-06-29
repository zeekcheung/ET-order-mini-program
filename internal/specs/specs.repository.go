package specs

import (
	"ET-order-mini-program/database/models"
	"errors"

	"gorm.io/gorm"
)

type SpecsRepository struct {
	db *gorm.DB
}

func NewSpecsRepository(db *gorm.DB) *SpecsRepository {
	return &SpecsRepository{db: db}
}

func (r *SpecsRepository) Create(specs *models.GoodsSpecs) error {
	result := r.db.Create(specs)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to create specs")
	}
	return nil
}

func (r *SpecsRepository) Update(specs *models.GoodsSpecs) error {
	result := r.db.Save(specs)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update specs")
	}
	return nil
}

func (r *SpecsRepository) Delete(specs *models.GoodsSpecs) error {
	result := r.db.Delete(specs)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete goods")
	}
	return nil
}

func (r *SpecsRepository) GetById(id uint) (*models.GoodsSpecs, error) {
	var specs models.GoodsSpecs
	result := r.db.First(&specs, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("goods not found")
	}
	return &specs, nil
}

func (r *SpecsRepository) GetAll() ([]models.GoodsSpecs, error) {
	var specs []models.GoodsSpecs
	result := r.db.Find(&specs)
	if result.Error != nil {
		return nil, result.Error
	}
	return specs, nil
}
