package dao

import (
	"fmt"
	"gg/models"
)

// GetUser for dao.
func GetUser() *models.UserInfo {
	userInfo := new(models.UserInfo)
	result := db.First(userInfo)
	fmt.Println("rowsaffected: ", result.RowsAffected)
	return userInfo
}

// GetUserInfoByName return UserInfo whose name is equal to the input name.
func GetUserInfoByName(name string) *models.UserInfo {
	userInfo := new(models.UserInfo)
	db.Find(userInfo, "Name=?", name)
	return userInfo
}
