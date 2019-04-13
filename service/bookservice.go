package service

import (
	"database/sql"

	"github.com/go-web-app/models"
	"github.com/go-web-app/repository"
)

type BookService struct {
	repo repository.BookRepo
}

func NewBookService(db *sql.DB) BookService {
	return BookService{
		repo: repository.NewBookRepo(db),
	}
}

func (bs *BookService) SaveBook(book models.Book) int {
	bs.repo.Save(book)
	return 1
}
