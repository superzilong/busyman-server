package service

import (
	"fmt"
	"gg/models"
)

// GetUser for dao.
func GetUser() *models.User {
	user := new(models.User)
	result := db.First(user)
	fmt.Println("rowsaffected: ", result.RowsAffected)
	return user
}

// GetUserByName return User whose name is equal to the input name.
func GetUserByName(name string) *models.User {
	user := new(models.User)
	result := db.Find(user, "name = ?", name)
	if result.Error == nil {
		return user
	}
	return nil
}
