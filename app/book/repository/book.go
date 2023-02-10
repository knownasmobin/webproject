package repository

import (
	"context"
	"strconv"
	"git.ecobin.ir/ecomicro/template/app/book/domain"

	"git.ecobin.ir/ecomicro/tooty"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bookRepository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &bookRepository{}

func NewBookRepository(dbConnection *gorm.DB) *bookRepository {
	err := dbConnection.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}
	return &bookRepository{dbConnection}
}

func (ur *bookRepository) Create(ctx context.Context, domainBook domain.Book) (*domain.Book, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] create book", "repository")
	defer tooty.CloseTheAPMSpan(span)

	bookDao := FromDomainBook(domainBook)
	result := ur.Conn.Debug().Create(&bookDao)
	if result.Error != nil {
		return nil, result.Error
	}
	book := bookDao.ToDomainBook()
	return &book, nil
}

func (ur *bookRepository) GetByCategory(ctx context.Context, categoryId *int) ([]domain.Book, error) {
	var bookArray []Book
	chain := ur.Conn.WithContext(ctx).Debug()
	if strconv.Itoa(*categoryId) != "0" {
		chain = chain.Where("categories->>'item' = ?", strconv.Itoa(*categoryId))
	} else {
                chain = chain.Find(&bookArray)
        }
	err := chain.Find(&bookArray).Error
	if err != nil {
		return nil, err
	}
	domainBooks := make([]domain.Book, len(bookArray))
	for idx, book := range bookArray {
		domainBooks[idx] = book.ToDomainBook()
	}
	return domainBooks, nil
}
func (ur *bookRepository) GetBookById(ctx context.Context, id int) (*domain.Book, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] get book by id", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var bookDao Book
	err := ur.Conn.WithContext(ctx).Debug().Where(Book{Id: id}).First(&bookDao).Error
	if err != nil {
		return nil, err
	}
	book := bookDao.ToDomainBook()
	return &book, nil
}

func (ur *bookRepository) Update(ctx context.Context, condition domain.Book, domainBook domain.Book) ([]domain.Book, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[R] update book", "repository")
	defer tooty.CloseTheAPMSpan(span)
	var bookArray []Book
	err := ur.Conn.WithContext(ctx).Debug().Model(&bookArray).Clauses(clause.Returning{}).Where(FromDomainBook(condition)).Updates(FromDomainBook(domainBook)).Error
	if err != nil {
		return []domain.Book{}, err
	}
	domainBooks := make([]domain.Book, len(bookArray))
	for idx, book := range bookArray {
		domainBooks[idx] = book.ToDomainBook()
	}
	return domainBooks, nil
}
