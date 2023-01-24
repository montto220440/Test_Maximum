package models

import "gorm.io/gorm"

type Todo_list struct {
	gorm.Model
	Content string `gorm:"uniqueIndex;type:varchar(100);not null"`
}