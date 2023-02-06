package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/template/app/category/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	Usecase domain.Usecase
}

func NewCategoryHandler(g *gin.Engine, authMiddleware gin.HandlerFunc, uu domain.Usecase) {
	rg := g.Group("/category")
	handler := &categoryHandler{
		Usecase: uu,
	}
	rg.GET("/",
		//  authMiddleware,
		handler.getCategory)
	rg.GET("/:id",
		//  authMiddleware,
		handler.getCategoryById)
	rg.POST("",
		//  authMiddleware,
		handler.createCategory)
	rg.PUT("",
		//  authMiddleware,
		handler.updateCategory)
}

// get category godoc
// @Summary get category
// @Schemes
// @Description  get category
// @Tags category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {object} domain.Category
// @Router /category [get]
func (uh *categoryHandler) getCategory(ctx *gin.Context) {
	categoryId, err := strconv.ParseUint(ctx.GetString("categoryId"), 10, 64)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	category, err := uh.Usecase.GetCategoryById(ctx, categoryId)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// get category by id godoc
// @Summary get category by id
// @Schemes
// @Description  get category by id
// @Tags category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path uint64 true "category id"
// @Success 200 {object} domain.Category
// @Router /category/{id} [get]
func (uh *categoryHandler) getCategoryById(ctx *gin.Context) {

	var uri CategoryIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	category, err := uh.Usecase.GetCategoryById(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// create category godoc
// @Summary create category
// @Schemes
// @Description  create category
// @Tags category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body CreateCategoryBody true "body params"
// @Success 200 {object} domain.Category
// @Router /category [post]
func (uh *categoryHandler) createCategory(ctx *gin.Context) {
	var body CreateCategoryBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	category, err := uh.Usecase.Create(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// update category godoc
// @Summary update category
// @Schemes
// @Description  update category
// @Tags category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdateCategoryBody true "body params"
// @Success 200 {object} domain.Category
// @Router /category [put]
func (uh *categoryHandler) updateCategory(ctx *gin.Context) {
	var body UpdateCategoryBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	category, err := uh.Usecase.Update(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, category)
}
