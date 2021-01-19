package models

import (
	"time"

	"gorm.io/gorm"
)

// Model 基类
type Model struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
