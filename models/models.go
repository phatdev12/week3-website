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
	ImageUrl   string  `gorm:"type:text;not null" json:"image_url" form:"image_url"`
	Price      float64 `gorm:"type:decimal(20,2);not null" json:"price" form:"price"`
	Quantity   int     `gorm:"not null"`
	IDCategory uint    `gorm:"not null" json:"id_category" form:"id_category"`
}
