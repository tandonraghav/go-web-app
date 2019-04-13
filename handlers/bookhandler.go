package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-web-app/configs"
	"github.com/go-web-app/models"
	"github.com/gorilla/mux"

	"github.com/go-web-app/service"
)

type BookHandler struct {
	bs service.BookService
}

func RegisterBookHandlers(db *configs.DB, r *mux.Router) {
	bookHandler := NewBookHandler(db.SQL)
	r.HandleFunc("/hello/{id}", bookHandler.SaveBookHandler).Methods("GET")
}

func NewBookHandler(db *sql.DB) BookHandler {
	return BookHandler{
		bs: service.NewBookService(db),
	}
}

func (bh *BookHandler) SaveBookHandler(w http.ResponseWriter, r *http.Request) {
	id := bh.bs.SaveBook(models.Book{ID: 1})
	log.Println("Id = %s of the book", id)
}
