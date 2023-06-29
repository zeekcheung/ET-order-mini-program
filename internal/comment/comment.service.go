package comment

import "ET-order-mini-program/database/models"

type CommentService struct {
	repo *CommentRepository
}

func NewCommentService(repo *CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Createcomment(comment *models.GoodsComment) error {
	return s.repo.Create(comment)
}

func (s *CommentService) GetcommentById(id uint) (*models.GoodsComment, error) {
	return s.repo.GetById(id)
}

func (s *CommentService) GetAllcomment() ([]models.GoodsComment, error) {
	return s.repo.GetAll()
}

func (s *CommentService) Updatecomment(comment *models.GoodsComment) error {
	return s.repo.Update(comment)
}

func (s *CommentService) Deletecomment(id uint) error {
	comment, _ := s.GetcommentById(id)
	return s.repo.Delete(comment)
}
