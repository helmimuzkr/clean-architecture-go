package repository

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	NPM  int    `gorm:"column:npm"`
	Name string `gorm:"column:name"`
}
