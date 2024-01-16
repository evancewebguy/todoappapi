package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID   string `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	gorm.Model
}
