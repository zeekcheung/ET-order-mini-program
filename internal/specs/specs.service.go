package specs

import "ET-order-mini-program/database/models"

type SpecsService struct {
	repo *SpecsRepository
}

func NewSpecsService(repo *SpecsRepository) *SpecsService {
	return &SpecsService{repo: repo}
}

func (s *SpecsService) CreateSpecs(specs *models.GoodsSpecs) error {
	return s.repo.Create(specs)
}

func (s *SpecsService) GetSpecsById(id uint) (*models.GoodsSpecs, error) {
	return s.repo.GetById(id)
}

func (s *SpecsService) GetAllSpecs() ([]models.GoodsSpecs, error) {
	return s.repo.GetAll()
}

func (s *SpecsService) UpdateSpecs(specs *models.GoodsSpecs) error {
	return s.repo.Update(specs)
}

func (s *SpecsService) DeleteSpecs(id uint) error {
	specs, _ := s.GetSpecsById(id)
	return s.repo.Delete(specs)
}
