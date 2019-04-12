package repository

import (
	"github.com/go-web-app/models"
)

type bookRepoImpl struct {
}

func NewBookRepo() BookRepo {
	return &bookRepoImpl{}
}
func (b *bookRepoImpl) FetchByID(id int) (*models.Book, error) {
	return nil, nil
}

func (b *bookRepoImpl) Save(book models.Book) (int, error) {
	return 1, nil
}
