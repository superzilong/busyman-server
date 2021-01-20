package service

import (
	"database/sql"
	"fmt"
	"gg/models"
	"gg/pkg/settings"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql" // Initialize mysql driver
	"gorm.io/gorm"
)

var db *gorm.DB
var mysqldb *sql.DB

// Init gorm.
func Init() (err error) {
	config := settings.Conf.MySQLConfig
	var source = fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DB,
	)
	//fmt.Printf("%+v/n", source)
	//db, err := gorm.Open("mysql", "", "root:1@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	db, err = gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqldb, err = db.DB()
	if err != nil {
		panic(err)
	}
	mysqldb.SetMaxIdleConns(config.MaxIdleConns)
	mysqldb.SetMaxOpenConns(config.MaxOpenConns)

	// 自动迁移
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Customer{},
		&models.SalesOrder{},
		&models.SalesOrderItem{},
		&models.Vendor{},
		&models.PurchaseOrder{},
		&models.PurchaseOrderItem{},
	)

	return
}

// Close gorm.
func Close() {
	mysqldb.Close()
}
