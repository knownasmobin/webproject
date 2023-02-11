package repository

import (
	"time"

	"git.ecobin.ir/ecomicro/template/app/comment/domain"
	"gorm.io/gorm"
)

type Comment struct {
	Id        int `gorm:"primaryKey;unique"`
	UserId    int
	BookId    int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainComment(comment domain.Comment) Comment {
	return Comment{
		Id:        comment.Id,
		UserId:    comment.UserId,
		BookId:    comment.BookId,
		Message:   comment.Message,
		CreatedAt: comment.CreatedDate,
		UpdatedAt: comment.UpdatedDate,
	}
}

func (u *Comment) ToDomainComment() domain.Comment {
	return domain.Comment{
		Id:          u.Id,
		UserId:      u.UserId,
		BookId:      u.BookId,
		Message:     u.Message,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
	}
}
