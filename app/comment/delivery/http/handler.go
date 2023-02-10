package http

import (
	"log"
	"net/http"

	"git.ecobin.ir/ecomicro/template/app/comment/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	Usecase domain.Usecase
}

func NewCommentHandler(g *gin.Engine, uu domain.Usecase) {
	rg := g.Group("/comment")
	handler := &commentHandler{
		Usecase: uu,
	}

	rg.PUT("/byCondition",
		//  authMiddleware,
		handler.getByCondition)
	rg.DELETE("/:id",
		//  authMiddleware,
		handler.delete)
	rg.GET("/:id",
		//  authMiddleware,
		handler.getCommentById)
	rg.POST("",
		//  authMiddleware,
		handler.createComment)
	rg.PUT("",
		//  authMiddleware,
		handler.updateComment)
}

// get comment by condition godoc
// @Summary get comment by condition
// @Schemes
// @Description  get comment by condition
// @Tags comment
// @Accept json
// @Produce json
// @Param body body UpdateCommentBody true "body params"
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {array} domain.Comment
// @Router /comment/byCondition [put]
func (uh *commentHandler) getByCondition(ctx *gin.Context) {
	var body UpdateCommentBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	comments, err := uh.Usecase.GetByCondition(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

// get comment by id godoc
// @Summary get comment by id
// @Schemes
// @Description  get comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path int true "comment id"
// @Success 200 {object} domain.Comment
// @Router /comment/{id} [get]
func (uh *commentHandler) getCommentById(ctx *gin.Context) {

	var uri CommentIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	comment, err := uh.Usecase.GetCommentById(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

// delete comment by id godoc
// @Summary delete comment by id
// @Schemes
// @Description  delete comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path int true "comment id"
// @Success 200 {object} domain.Comment
// @Router /comment/{id} [delete]
func (uh *commentHandler) delete(ctx *gin.Context) {
	var uri CommentIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	comment, err := uh.Usecase.Delete(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

// create comment godoc
// @Summary create comment
// @Schemes
// @Description  create comment
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body CreateCommentBody true "body params"
// @Success 200 {object} domain.Comment
// @Router /comment [post]
func (uh *commentHandler) createComment(ctx *gin.Context) {
	var body CreateCommentBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	comment, err := uh.Usecase.Create(ctx, body.toDomain())
	print("commented")
	log.Println(comment)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

// update comment godoc
// @Summary update comment
// @Schemes
// @Description  update comment
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdateCommentBody true "body params"
// @Success 200 {object} domain.Comment
// @Router /comment [put]
func (uh *commentHandler) updateComment(ctx *gin.Context) {
	var body UpdateCommentBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	comment, err := uh.Usecase.Update(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
