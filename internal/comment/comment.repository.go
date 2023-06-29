package comment

import (
	"ET-order-mini-program/database/models"
	"errors"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *models.GoodsComment) error {
	result := r.db.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to create goods")
	}
	return nil
}

func (r *CommentRepository) Update(comment *models.GoodsComment) error {
	result := r.db.Save(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update goods")
	}
	return nil
}

func (r *CommentRepository) Delete(comment *models.GoodsComment) error {
	result := r.db.Delete(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete goods")
	}
	return nil
}

func (r *CommentRepository) GetById(id uint) (*models.GoodsComment, error) {
	var comment models.GoodsComment
	result := r.db.First(&comment, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("goods not found")
	}
	return &comment, nil
}

func (r *CommentRepository) GetAll() ([]models.GoodsComment, error) {
	var comments []models.GoodsComment
	result := r.db.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
