package models

import (
	"github.com/jinzhu/gorm"
	"github.com/turboenot/simple-REST/pkg/config"
)

var db *gorm.DB

type Car struct {
	gorm.Model
	//Id          string `json:"id"`
	Brand   string `gorm:""json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  string `json:"status"`
	Mileage uint32 `json:"mileage"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Car{})
}

func (b *Car) CreateCar() *Car {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllCars() []Car {
	var Cars []Car
	db.Find(&Cars)
	return Cars
}

func GetCarById(Id int64) (*Car, *gorm.DB) {
	var getCar Car
	db := db.Where("ID = ?", Id).Find(&getCar)
	return &getCar, db
}

func DeleteCar(ID int64) Car {
	var car Car
	db.Where("ID = ?", ID).Delete(car)
	return car
}
