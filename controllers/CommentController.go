package controllers

import (
	"encoding/json"
	"go-news-api/database"
	"go-news-api/entities"
	"net/http"
)

func CreateComments(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var comment entities.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	var news *entities.News
	database.Instance.First(&news, comment.NewsID)
	if news.ID == 0{
		json.NewEncoder(w).Encode("News not found!")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	comment.News = *news

	database.Instance.Create(&comment)
	json.NewEncoder(w).Encode(comment)
}

func GetComments(w http.ResponseWriter, r *http.Request){
	var comments []entities.Comment
	database.Instance.Find(&comments)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}