package routes

import (
	"github.com/SeanardK/GO-REST_API/pkg/controller"
	"github.com/gorilla/mux"
)

func IndexRoutes(Router *mux.Router) {
	Router.HandleFunc("/index", controller.IndexGetAll).Methods("GET")
}
