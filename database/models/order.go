package models

type Order struct {
	Model

	Status         int    `json:"status"`
	Remarks        string `json:"remarks"`
	PackingCharges int    `json:"packingCharges"`
	PaymentMethod  int    `jsoN:"paymentMethod"`
	TotalPrice     int    `json:"totalPrice"`

	// 用户与订单的一对多关系
	UserID string `json:"userId"`

	// 订单与订单列表项的一对多关系
	Items []OrderItem `json:"items"`

	// 订单与收货地址的多对一关系
	AddressID uint `json:"addressId"`
}
