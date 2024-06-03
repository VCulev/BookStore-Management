package controllers

import (
	"net/http"
)

func HandleErrors(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func GetBook(w http.ResponseWriter, r *http.Request) {

}

func GetBookById(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
