package models

// UserInfo 用户信息
type UserInfo struct {
	ID       uint
	Name     string `json:"Name" binding:"required"`
	Password string `json:"Password" binding:"required"`
	Gender   string
	Hobby    string
}
