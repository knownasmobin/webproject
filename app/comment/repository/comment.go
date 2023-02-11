package repository

import (
	"context"
	"log"

	"git.ecobin.ir/ecomicro/template/app/comment/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepository struct {
	Conn *gorm.DB
}

var _ domain.Repository = &commentRepository{}

func NewCommentRepository(dbConnection *gorm.DB) *commentRepository {
	err := dbConnection.AutoMigrate(&Comment{})
	if err != nil {
		panic(err)
	}
	return &commentRepository{dbConnection}
}

func (ur *commentRepository) Create(ctx context.Context, domainComment domain.Comment) (*domain.Comment, error) {
	log.Println("kir khare ripo")
	commentDao := FromDomainComment(domainComment)
	result := ur.Conn.Debug().Create(&commentDao)
	if result.Error != nil {
		return nil, result.Error
	}
	comment := commentDao.ToDomainComment()
	return &comment, nil
}
func (ur *commentRepository) Delete(ctx context.Context, id int) (*domain.Comment, error) {
	var commentArray []Comment
	result := ur.Conn.Debug().Delete(&commentArray, Comment{Id: id})
	if result.Error != nil {
		return nil, result.Error
	}
	var comment *domain.Comment = nil
	if len(commentArray) > 0 {
		temp := commentArray[0].ToDomainComment()
		comment = &temp
	}
	return comment, nil
}
func (ur *commentRepository) GetCommentById(ctx context.Context, id int) (*domain.Comment, error) {
	var commentDao Comment
	err := ur.Conn.WithContext(ctx).Debug().Where(Comment{Id: id}).Find(&commentDao).Error
	if err != nil {
		return nil, err
	}
	comment := commentDao.ToDomainComment()
	return &comment, nil
}
func (ur *commentRepository) GetByCondition(ctx context.Context, condition domain.Comment) ([]domain.Comment, error) {
	var commentArray []Comment
	err := ur.Conn.WithContext(ctx).Debug().Where(FromDomainComment(condition)).Find(&commentArray).Error
	if err != nil {
		return []domain.Comment{}, err
	}
	domainComments := make([]domain.Comment, len(commentArray))
	for idx, comment := range commentArray {
		domainComments[idx] = comment.ToDomainComment()
	}
	return domainComments, nil
}

func (ur *commentRepository) Update(ctx context.Context, condition domain.Comment, domainComment domain.Comment) ([]domain.Comment, error) {

	var commentArray []Comment
	err := ur.Conn.WithContext(ctx).Debug().Model(&commentArray).Clauses(clause.Returning{}).Where(FromDomainComment(condition)).Updates(FromDomainComment(domainComment)).Error
	if err != nil {
		return []domain.Comment{}, err
	}
	domainComments := make([]domain.Comment, len(commentArray))
	for idx, comment := range commentArray {
		domainComments[idx] = comment.ToDomainComment()
	}
	return domainComments, nil
}
