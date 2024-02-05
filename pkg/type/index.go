package types

import "github.com/SeanardK/GO-REST_API/pkg/models"

type Response struct {
	Status     string        `json:"message"`
	StatucCode int           `json:"status"`
	Data       []models.Book `json:"data"`
}
