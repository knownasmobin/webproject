package http

import "git.ecobin.ir/ecomicro/template/app/comment/domain"

type CreateCommentBody struct {
}
type UpdateCommentBody struct {
	Id uint64 `json:"id"`
}

func (c UpdateCommentBody) toDomain() domain.Comment {
	return domain.Comment{
		Id: c.Id,
	}
}

func (c CreateCommentBody) toDomain() domain.Comment {
	return domain.Comment{}
}

type CommentIdUri struct {
	Id uint64 `uri:"id"`
}
