package routes

import (
	"github.com/SeanardK/GO-REST_API/pkg/controller"
	"github.com/gorilla/mux"
)

func BookRoutes(Router *mux.Router) {
	Router.HandleFunc("/books", controller.BookCreate).Methods("POST")
	Router.HandleFunc("/books", controller.BookGetAll).Methods("GET")
}
