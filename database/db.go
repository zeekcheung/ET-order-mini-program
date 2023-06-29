package database

import (
	"ET-order-mini-program/configs"
	"ET-order-mini-program/database/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func InitDB(config *configs.Config) (db *gorm.DB) {
	// 连接数据库
	db = connectDB(config)

	// 自动建表
	if err := db.AutoMigrate(
		&models.Address{},
		&models.CartItem{},
		&models.Cart{},
		&models.GoodsCategory{},
		&models.GoodsComment{},
		&models.GoodsSpecs{},
		&models.Goods{},
		&models.OrderItem{},
		&models.Order{},
		&models.Shop{},
	); err != nil {
		panic("Failed to auto migrate, error: " + err.Error())
	}

	return db
}

// 连接数据库
func connectDB(config *configs.Config) (db *gorm.DB) {
	dbConfig := config.Database

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.Charset,
		dbConfig.Location,
	)

	gormConfig := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)

	if err != nil {
		panic("Failed to connect database, error: " + err.Error())
	}

	return db
}
