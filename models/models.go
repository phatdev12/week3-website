package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
}

type Product struct {
	gorm.Model
	Id         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"type:varchar(100);not null"`
	Price      float64 `gorm:"type:decimal(20,2);not null"`
	Quantity   int     `gorm:"not null"`
	IDCategory uint    `gorm:"not null"`
}
