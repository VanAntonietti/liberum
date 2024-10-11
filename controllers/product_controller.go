package controllers

import (
	"encoding/json"
	"net/http"

	"liberum/database"
	"liberum/models"
)

func GetProducts(response http.ResponseWriter, request *http.Request) {
	var products []models.Product
	database.DB.Find(&products)

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(products)
}

func CreateProducts(response http.ResponseWriter, request *http.Request) {
	var product models.Product
	json.NewDecoder(request.Body).Decode(&product)

	database.DB.Create(&product)

	response.Header().Set("Content Type", "application/json")
	json.NewEncoder(response).Encode(product)
}