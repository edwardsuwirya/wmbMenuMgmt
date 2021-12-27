package entity

import "gorm.io/gorm"

type Menu struct {
	ID       string `gorm:"size:5;primaryKey"`
	MenuName string `gorm:"size:100"`

	Price int
	gorm.Model
}

func (c *Menu) TableName() string {
	return "menu"
}
