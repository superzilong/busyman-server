package models

// User 用户
type User struct {
	Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
