package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func GetNewsById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newsId := mux.Vars(r)["id"]
	if checkIfNewsExists(newsId) == false{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "News not found"})
		return
	}
	var news entities.News
	database.Instance.First(&news, newsId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}

func DeleteNews(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newsId := mux.Vars(r)["id"]
	if checkIfNewsExists(newsId) == false{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "News not found"})
		return
	}
	var news entities.News
	database.Instance.Delete(&news, newsId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(news)
}

func checkIfNewsExists(newsId string) bool{
	var news entities.News
	database.Instance.First(&news, newsId)
	if news.ID == 0{
		return false
	}
	return true
}