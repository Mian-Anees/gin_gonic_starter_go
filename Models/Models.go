package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name      string `gorm:"column:name"  json:"name"`
	Type      string `gorm:"column:type"  json:"type"`
	Available bool   `gorm:"column:available"  json:"available"`
}

type Todo struct {
	ID          uint   `gorm:"column:id"  json:"id"`
	Title       string `gorm:"column:title"  json:"title"`
	Description string `gorm:"column:description"  json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}
