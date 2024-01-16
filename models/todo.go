package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	Id        string `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	gorm.Model
}
