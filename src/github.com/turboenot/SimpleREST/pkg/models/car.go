package models

import (
	"github.com/example/simple-REST/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Car struct {
	gorm.Model
	//Id          string `json:"id"`
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
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
