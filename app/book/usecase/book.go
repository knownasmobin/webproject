package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/book/domain"
)

type bookUsecase struct {
	bookRepo domain.Repository
}

var _ domain.Usecase = &bookUsecase{}
var _ domain.Adapter = &bookUsecase{}

func NewBookUsecase(bookRepo domain.Repository) *bookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}
func (uu *bookUsecase) SetAdapters() {

}
func (uu *bookUsecase) Create(
	ctx context.Context,
	book domain.Book,
) (*domain.Book, error) {

	dbBook, err := uu.bookRepo.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	return dbBook, nil
}

func (uu *bookUsecase) Update(ctx context.Context, book domain.Book) (*domain.Book, error) {
	bookArray, err := uu.bookRepo.Update(ctx, domain.Book{
		Id: book.Id,
	}, book)
	if err != nil {
		return nil, err
	}
	if len(bookArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &bookArray[0], nil
}
func (uu *bookUsecase) GetAll(ctx context.Context, categoryId *int) ([]domain.Book, error) {
	books, err := uu.bookRepo.GetByCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (uu *bookUsecase) GetBookById(ctx context.Context, id int) (*domain.Book, error) {
	book, err := uu.bookRepo.GetBookById(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (uu *bookUsecase) GetByCondition(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	dbBook, err := uu.bookRepo.GetByCondition(ctx, book)
	if err != nil {
		return nil, err
	}
	return dbBook, nil
}
