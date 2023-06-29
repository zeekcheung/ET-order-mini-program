package goods

import (
	"ET-order-mini-program/database/models"
)

type CreateGoodsDto struct {
	Name           string
	ImageUrl       string
	Description    string
	IsExchangeable bool
	ExchangeLimit  int
	ExchangeCost   int

	ShopID     uint
	GoodsSpecs []struct {
		Name        string
		Description string
		Price       float32
	}
}

type UpdateGoodsDto CreateGoodsDto

func createGoodsFromDto(d CreateGoodsDto) *models.Goods {
	// 创建商品规格
	var goodsSpecs []models.GoodsSpecs
	for _, specs := range d.GoodsSpecs {
		goodsSpecs = append(goodsSpecs, models.GoodsSpecs{
			Name:        specs.Name,
			Description: d.Description,
			Price:       specs.Price,
		})
	}

	return &models.Goods{
		Name:           d.Name,
		ImageUrl:       d.ImageUrl,
		Description:    d.Description,
		IsExchangeable: d.IsExchangeable,
		ExchangeCost:   d.ExchangeCost,
		ExchangeLimit:  d.ExchangeLimit,
		GoodsSpecs:     goodsSpecs,
	}
}

func (d CreateGoodsDto) CreateGoods() *models.Goods {
	// 创建商品规格
	var goodsSpecs []models.GoodsSpecs
	for _, specs := range d.GoodsSpecs {
		goodsSpecs = append(goodsSpecs, models.GoodsSpecs{
			Name:        specs.Name,
			Description: d.Description,
			Price:       specs.Price,
		})
	}

	return &models.Goods{
		Name:           d.Name,
		ImageUrl:       d.ImageUrl,
		Description:    d.Description,
		IsExchangeable: d.IsExchangeable,
		ExchangeCost:   d.ExchangeCost,
		ExchangeLimit:  d.ExchangeLimit,
		GoodsSpecs:     goodsSpecs,
	}
}

func (d UpdateGoodsDto) CreateGoods() *models.Goods {
	return createGoodsFromDto(CreateGoodsDto(d))
}
