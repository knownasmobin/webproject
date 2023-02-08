package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/template/app/book/domain"
	"git.ecobin.ir/ecomicro/template/common"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	Usecase domain.Usecase
}

func NewBookHandler(g *gin.Engine, uu domain.Usecase) {
	rg := g.Group("/book")
	handler := &bookHandler{
		Usecase: uu,
	}
	rg.GET("/",
		common.GetAuth().Guard(true),
		handler.getAll)
	rg.GET("/:id",
		common.GetAuth().Guard(false),
		handler.getBookById)
	rg.POST("",
		//  authMiddleware,
		handler.createBook)
	rg.PUT("",
		//  authMiddleware,
		handler.updateBook)
}

// get book godoc
// @Summary get book
// @Schemes
// @Description  get book
// @Tags book
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param categoryId query int false "category id"
// @Success 200 {array} domain.Book
// @Router /book [get]
func (uh *bookHandler) getAll(ctx *gin.Context) {
	var query IdFromQuery
	err := ctx.Bind(&query)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	books, err := uh.Usecase.GetAll(ctx, &query.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

// get book by id godoc
// @Summary get book by id
// @Schemes
// @Description  get book by id
// @Tags book
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path int true "book id"
// @Success 200 {object} domain.Book
// @Router /book/{id} [get]
func (uh *bookHandler) getBookById(ctx *gin.Context) {

	var uri BookIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	book, err := uh.Usecase.GetBookById(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// create book godoc
// @Summary create book
// @Schemes
// @Description  create book
// @Tags book
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body CreateBookBody true "body params"
// @Success 200 {object} domain.Book
// @Router /book [post]
func (uh *bookHandler) createBook(ctx *gin.Context) {
	var body CreateBookBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	book, err := uh.Usecase.Create(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// update book godoc
// @Summary update book
// @Schemes
// @Description  update book
// @Tags book
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdateBookBody true "body params"
// @Success 200 {object} domain.Book
// @Router /book [put]
func (uh *bookHandler) updateBook(ctx *gin.Context) {
	var body UpdateBookBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	book, err := uh.Usecase.Update(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, book)
}
