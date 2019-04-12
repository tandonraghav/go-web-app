package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-web-app/models"
	"github.com/go-web-app/repository"
	"github.com/gorilla/mux"
)

type BookRepo struct {
	repo repository.BookRepo
}

var bookRepo = BookRepo{
	repo: repository.NewBookRepo(),
}

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{id}", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("param = %s", vars)
	fmt.Fprintf(w, "Hello World! ")
	bookRepo.repo.Save(models.Book{ID: 10})
}
