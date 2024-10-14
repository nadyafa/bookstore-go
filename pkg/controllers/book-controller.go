package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nadyafa/bookstore-apps/pkg/models"
	"github.com/nadyafa/bookstore-apps/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := models.Book{}
	utils.ParseBody(r, &updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(id)
	if updateBook.Title != "" { //try with nil next time
		bookDetails.Title = updateBook.Title
	}

	if updateBook.Author != "" { //try with nil next time
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" { //try with nil next time
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
