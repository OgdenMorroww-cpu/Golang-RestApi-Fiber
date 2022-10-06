package entities

import "gorm.io/gorm"

type CarProducts struct {
	gorm.Model
	Name         string `json:"name"`
	CarModel     string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Country      string `json:"country"`
	Price        uint   `json:"price"`
	IsBrandNew   bool   `json:"isBrandNew" gorm:"default:true"`
}
