package models

type Cart struct {
	Model

	// 用户与购物车的一对一关系
	UserID string `json:"userId"`

	// 购物车与购物车列表项的一对多关系
	Items []CartItem `json:"items"`
}
