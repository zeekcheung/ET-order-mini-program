package goods

import (
	"ET-order-mini-program/database/models"
	"ET-order-mini-program/internal/shop"
	"ET-order-mini-program/internal/specs"

	"gorm.io/gorm"
)

type GoodsService struct {
	*gorm.DB
	goodsRepo *GoodsRepository
	shopRepo  *shop.ShopRepository
	specsRepo *specs.SpecsRepository
}

func NewGoodsService(db *gorm.DB) *GoodsService {
	return &GoodsService{
		DB:        db,
		goodsRepo: NewGoodsRepository(db),
		shopRepo:  shop.NewShopRepository(db),
		specsRepo: specs.NewSpecsRepository(db),
	}
}

func (s *GoodsService) Creategoods(dto *CreateGoodsDto) (*models.Goods, error) {
	// 加载商铺实体
	shop, err := s.shopRepo.GetById(dto.ShopID)
	if err != nil {
		return nil, err
	}
	// 创建商品实体
	goods := dto.CreateGoods()
	// 关联商铺实体
	goods.ShopID = shop.ID

	// 保存到数据库
	// Gorm 会自动创建关系对应的实体，不需要通过事务手动实现
	if err := s.goodsRepo.Create(goods); err != nil {
		return nil, err
	}

	return goods, nil
}

func (s *GoodsService) GetgoodsById(id string) (*models.Goods, error) {
	return s.goodsRepo.GetById(id)
}

func (s *GoodsService) GetAllGoods() ([]models.Goods, error) {
	return s.goodsRepo.GetAll()
}

func (s *GoodsService) UpdateGoods(goods *models.Goods) error {
	return s.goodsRepo.Update(goods)
}

func (s *GoodsService) DeleteGoods(id string) error {
	goods, _ := s.GetgoodsById(id)
	return s.goodsRepo.Delete(goods)
}
