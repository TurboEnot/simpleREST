package routes

import (
	"github.com/example/simple-REST/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCarStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/car/", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/car/", controllers.GetCar).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.GetCarById).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/car/{carId}", controllers.DeleteCar).Methods("DELETE")
}Car