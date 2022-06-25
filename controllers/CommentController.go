package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func DeleteComment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	commentId := mux.Vars(r)["id"]
	if checkIfCommentExists(commentId) == false{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Comment not found"})
		return
	}
	var comment entities.Comment
	database.Instance.Delete(&comment, commentId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Comment Deleted Successfully!")
}

func checkIfCommentExists(commentId string) bool{
	var comment entities.Comment
	database.Instance.First(&comment, commentId)
	if comment.ID == 0{
		return false
	}
	return true
}