package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SeanardK/GO-REST_API/pkg/models"
	types "github.com/SeanardK/GO-REST_API/pkg/type"
)

func BookCreate(w http.ResponseWriter, r *http.Request) {
	book := models.BookCreate(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&types.Response{"Success", 201, []models.Book{book}})
}

func BookGetAll(w http.ResponseWriter, r *http.Request) {
	books := models.BookGetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&types.Response{"success", 200, books})
}
