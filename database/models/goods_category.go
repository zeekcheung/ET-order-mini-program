package models

type GoodsCategory struct {
	Model

	Name        string `json:"name"`
	Description string `json:"description"`

	Goods []Goods `gorm:"many2many:goods_category;" json:"goods"`
}
