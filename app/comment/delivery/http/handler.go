package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/template/app/comment/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	Usecase domain.Usecase
}

func NewCommentHandler(g *gin.Engine, authMiddleware gin.HandlerFunc, uu domain.Usecase) {
	rg := g.Group("/comment")
	handler := &commentHandler{
		Usecase: uu,
	}
	rg.GET("/",
		//  authMiddleware,
		handler.getComment)
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

// get comment godoc
// @Summary get comment
// @Schemes
// @Description  get comment
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {object} domain.Comment
// @Router /comment [get]
func (uh *commentHandler) getComment(ctx *gin.Context) {
	commentId, err := strconv.ParseUint(ctx.GetString("commentId"), 10, 64)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	comment, err := uh.Usecase.GetCommentById(ctx, commentId)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

// get comment by id godoc
// @Summary get comment by id
// @Schemes
// @Description  get comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path uint64 true "comment id"
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
