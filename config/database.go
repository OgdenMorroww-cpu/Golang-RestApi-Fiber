package config

import (
	"modules/modules/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func ConnectDataBase() error {
	var err error
	DataBase, err = gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DataBase.AutoMigrate(&entities.CarProducts{})
	return nil
}
