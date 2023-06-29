package models

type Address struct {
	Model

	ReceiverName    string `json:"receiverName"`
	ReceiverSex     int    `json:"receiverSex"`
	ReceiverAddress string `json:"receiverAddress"`
	ReceiverPhone   string `json:"receiverPhone"`

	// 订单与收货地址的多对一关系
	Orders []Order `json:"orders"`

	// 用户与收货地址的一对多关系
	UserID string `json:"userId"`
}
