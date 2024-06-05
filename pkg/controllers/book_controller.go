package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VCulev/BookStore-Management/pkg/config"
	"github.com/VCulev/BookStore-Management/pkg/models"
	"github.com/VCulev/BookStore-Management/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	db := config.GetDB()
	db.Create(createBook)
	res, err := json.Marshal(createBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var books []models.Book
	db.Find(&books)
	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId:", err)
		http.Error(w, "Invalid bookId", http.StatusBadRequest)
		return
	}
	db := config.GetDB()
	var book models.Book
	result := db.First(&book, ID)
	if result.Error != nil {
		fmt.Println("Error while fetching book:", result.Error)
		http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId:", err)
		http.Error(w, "Invalid bookId", http.StatusBadRequest)
		return
	}
	db := config.GetDB()
	var book models.Book
	result := db.First(&book, ID)
	if result.Error != nil {
		fmt.Println("Error while fetching book:", result.Error)
		http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
		return
	}
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId:", err)
		http.Error(w, "Invalid bookId", http.StatusBadRequest)
		return
	}
	db := config.GetDB()
	var book models.Book
	result := db.Delete(&book, ID)
	if result.Error != nil {
		fmt.Println("Error while deleting book:", result.Error)
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
