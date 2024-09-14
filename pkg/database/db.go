package database

import (
	"gorm.io/gorm"
)

type CloudCost struct {
	Date    string
	Amount  float64
	Service string
}

func InsertCost(db *gorm.DB, cost CloudCost) error {
	return db.Create(&cost).Error
}
