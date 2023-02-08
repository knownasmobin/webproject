package usecase

import (
	"context"

	"git.ecobin.ir/ecomicro/template/app/comment/domain"
	"git.ecobin.ir/ecomicro/tooty"
)

type commentUsecase struct {
	commentRepo domain.Repository
}

var _ domain.Usecase = &commentUsecase{}
var _ domain.Adapter = &commentUsecase{}

func NewCommentUsecase(commentRepo domain.Repository) *commentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
	}
}
func (uu *commentUsecase) SetAdapters() {

}
func (uu *commentUsecase) Create(
	ctx context.Context,
	comment domain.Comment,
) (*domain.Comment, error) {

	span := tooty.OpenAnAPMSpan(ctx, "[U] create new comment", "usecase")
	defer tooty.CloseTheAPMSpan(span)

	dbComment, err := uu.commentRepo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return dbComment, nil
}
func (uu *commentUsecase) GetByCondition(ctx context.Context, comment domain.Comment) ([]domain.Comment, error) {

	comments, err := uu.commentRepo.GetByCondition(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (uu *commentUsecase) Update(ctx context.Context, comment domain.Comment) (*domain.Comment, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] update comment", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	commentArray, err := uu.commentRepo.Update(ctx, domain.Comment{
		Id: comment.Id,
	}, comment)
	if err != nil {
		return nil, err
	}
	if len(commentArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &commentArray[0], nil
}
func (uu *commentUsecase) GetCommentById(ctx context.Context, id int) (*domain.Comment, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get comment by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	comment, err := uu.commentRepo.GetCommentById(ctx, id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (uu *commentUsecase) Delete(ctx context.Context, id int) (*domain.Comment, error) {
	comment, err := uu.commentRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}