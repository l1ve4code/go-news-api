package controllers

import (
	"encoding/json"
	"go-news-api/database"
	"go-news-api/entities"
	"net/http"
)

func CreateNews(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var news entities.News
	json.NewDecoder(r.Body).Decode(&news)
	database.Instance.Create(&news)
	json.NewEncoder(w).Encode(news)
}

func GetNews(w http.ResponseWriter, r *http.Request){
	var news []entities.News
	database.Instance.Find(&news)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}