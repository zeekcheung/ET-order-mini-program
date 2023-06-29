package models

type GoodsSpecs struct {
	Model

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`

	// 商品与商品规格的一对多关系
	GoodsID uint `json:"goodsId"`

	// 购物车列表项与商品规格的多对一关系
	CartItems []CartItem `json:"cartItems"`

	// 订单列表项与商品规格的多对一关系
	OrderItems []OrderItem `json:"orderItems"`
}
