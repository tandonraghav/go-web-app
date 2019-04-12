package repository

import (
	"github.com/go-web-app/models"
)

type BookRepo interface {
	FetchByID(id int) (*models.Book, error)
	Save(book models.Book) (int, error)
}
