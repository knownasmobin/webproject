package http

import "git.ecobin.ir/ecomicro/template/app/comment/domain"

type CreateCommentBody struct {
	UserId  int
	BookId  int
	Message string
}
type UpdateCommentBody struct {
	Id      int
	UserId  int
	BookId  int
	Message string
}

func (c UpdateCommentBody) toDomain() domain.Comment {
	return domain.Comment{
		Id:      c.Id,
		UserId:  c.UserId,
		BookId:  c.BookId,
		Message: c.Message,
	}
}

func (c CreateCommentBody) toDomain() domain.Comment {
	return domain.Comment{
		UserId:  c.UserId,
		BookId:  c.BookId,
		Message: c.Message,
	}
}

type CommentIdUri struct {
	Id int `uri:"id"`
}
