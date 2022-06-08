package controllers

import (
	"encoding/json"
	"github.com/sean-cunniffe/go-bookstore/src/pkg/models"
	"github.com/sean-cunniffe/go-bookstore/src/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tempId := vars["id"]
	bookId, err := strconv.Atoi(tempId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book id"))
		return
	}
	book := models.GetBookById(bookId)
	emptyBook := models.Book{}
	if *book == emptyBook {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r.Body, &book)
	book = *book.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tempBookId := vars["id"]
	bookId, err := strconv.Atoi(tempBookId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book id"))
		return
	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tempBookId := vars["id"]
	bookId, err := strconv.Atoi(tempBookId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book id"))
		return
	}
	updateBook := &models.Book{}
	utils.ParseBody(r.Body, &updateBook)
	bookDetails := models.GetBookById(bookId)
	// should probably do some validating of data in body
	bookDetails.Author = updateBook.Author
	bookDetails.Name = updateBook.Name
	bookDetails.Publication = updateBook.Publication
	models.SaveBook(bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
