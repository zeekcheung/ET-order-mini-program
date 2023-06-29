package models

type GoodsComment struct {
	Model

	Content string `json:"content"`

	// 商品评论与用户的多对一关系
	UserID string `json:"userId"`

	// 商品与商品评论的一对多关系
	GoodsID string `json:"goodsId"`
}
