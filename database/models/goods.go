package models

type Goods struct {
	Model

	Name           string `json:"name"`
	ImageUrl       string `json:"imageUrl"`
	Description    string `json:"description"`
	IsExchangeable bool   `json:"isExchangeable"`
	ExchangeLimit  int    `json:"exchangeLimit"`
	ExchangeCost   int    `json:"exchangeCost"`

	// 商品与商品分类的多对多关系
	Categories []GoodsCategory `gorm:"many2many:goods_category;" json:"categories"`

	// 店铺与商品的一对多关系
	ShopID uint `json:"shopId"`

	// 商品与商品评论的一对多关系
	GoodsComments []GoodsComment `json:"goodsComments"`

	// 商品与商品规格的一对多关系
	GoodsSpecs []GoodsSpecs `json:"goodsSpecs"`

	// 购物车列表项与商品的多对一关系
	CartItems []CartItem `json:"cartItems"`

	// 订单列表项与商品的多对一关系
	OrderItems []OrderItem `json:"orderItems"`
}

func (g Goods) TableName() string {
	return "goods"
}
