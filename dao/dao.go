package dao

import (
	"fmt"
	"gg/models"
	"gg/pkg/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Initialize mysql driver
)

var db *gorm.DB

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
	db, err = gorm.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	// 自动迁移
	db.AutoMigrate(&models.UserInfo{})

	return
}

// Close gorm.
func Close() {
	db.Close()
}
