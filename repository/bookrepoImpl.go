package repository

import (
	"database/sql"

	"github.com/go-web-app/models"
)

type bookRepoImpl struct {
	Conn *sql.DB
}

func NewBookRepo(Conn *sql.DB) BookRepo {
	return &bookRepoImpl{
		Conn: Conn,
	}
}
func (b *bookRepoImpl) FetchByID(id int) (*models.Book, error) {
	return nil, nil
}

func (b *bookRepoImpl) Save(book models.Book) (int, error) {
	return 1, nil
}
