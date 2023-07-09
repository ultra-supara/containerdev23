package usecases

import (
	"context"

	"github.com/ultra-supara/containerdev23/pkg/entity"
)

type ListBooks struct {
	BookRepository entity.BookRepository
}

func NewListBooks(bookRepository entity.BookRepository) *ListBooks {
	return &ListBooks{BookRepository: bookRepository}
}

func (u *ListBooks) Execute(ctx context.Context) (*[]entity.Book, error) {
	books, err := u.BookRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
