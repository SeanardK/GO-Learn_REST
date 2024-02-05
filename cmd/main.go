package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SeanardK/GO-REST_API/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register Routes
	routes.BookRoutes(r)

	fmt.Print("Success Running App")

	log.Fatal(http.ListenAndServe("localhost:3001", r))
}
