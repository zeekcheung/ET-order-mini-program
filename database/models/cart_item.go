package models

type CartItem struct {
	Model

	Count int `json:"count"`

	// 购物车与购物车列表项的一对多关系
	CartID uint `json:"cartId"`

	// 购物车列表项与商品的多对一关系
	GoodsID uint `json:"goodsId"`

	// 购物车列表项与商品规格的多对一关系
	GoodsSpecsID uint `json:"goodsSpecsId"`
}
