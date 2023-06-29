package models

type Shop struct {
	Model

	OwnerID      string  `json:"ownerId"`
	Name         string  `json:"name"`
	AvatarUrl    string  `json:"avatarUrl"`
	Introduction string  `json:"introduction"`
	Grade        float32 `json:"grade"`
	Address      string  `json:"address"`

	// 商品与店铺的一对多关系
	Goods []Goods `json:"goods"`
}
