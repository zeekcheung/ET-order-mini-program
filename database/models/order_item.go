package models

type OrderItem struct {
	Model

	Count int `json:"count"`

	// 订单与订单列表项的一对多关系
	OrderID uint `json:"orderId"`

	// 订单列表项与商品的多对一关系
	GoodsID uint `json:"goodsId"`

	// 订单列表项与商品规格的多对一关系
	GoodsSpecsID uint `json:"goodsSpecsId"`
}
