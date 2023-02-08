package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type favoriteHandler struct {
	Usecase domain.Usecase
}

func NewFavoriteHandler(g *gin.Engine, uu domain.Usecase) {
	rg := g.Group("/favorite")
	handler := &favoriteHandler{
		Usecase: uu,
	}
	rg.GET("/mostFavorites",
		//  authMiddleware,
		handler.getMostFavorites)
	rg.PUT("/",
		//  authMiddleware,
		handler.getByCondition)
	rg.POST("",
		//  authMiddleware,
		handler.createFavorite)
	rg.DELETE("",
		//  authMiddleware,
		handler.deleteFavorite)
}

// get most favorite godoc
// @Summary get most favorite
// @Schemes
// @Description  get most favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {array} domain.Favorite
// @Router /favorite [get]
func (uh *favoriteHandler) getMostFavorites(ctx *gin.Context) {
	favorite, err := uh.Usecase.GetMostFavorites(ctx)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

// get by condition favorite godoc
// @Summary get by condition favorite
// @Schemes
// @Description  get by condition favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdateFavoriteBody true "body params"
// @Success 200 {array} domain.Favorite
// @Router /favorite [put]
func (uh *favoriteHandler) getByCondition(ctx *gin.Context) {
	var body UpdateFavoriteBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	favorite, err := uh.Usecase.GetByCondition(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

// create favorite godoc
// @Summary create favorite
// @Schemes
// @Description  create favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body CreateFavoriteBody true "body params"
// @Success 200 {object} domain.Favorite
// @Router /favorite [post]
func (uh *favoriteHandler) createFavorite(ctx *gin.Context) {
	var body CreateFavoriteBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	favorite, err := uh.Usecase.Create(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

// update favorite godoc
// @Summary update favorite
// @Schemes
// @Description  update favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id int true " favorite id"
// @Success 200 {object} domain.Favorite
// @Router /favorite/{id} [delete]
func (uh *favoriteHandler) deleteFavorite(ctx *gin.Context) {
	var uri FavoriteIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	favorite, err := uh.Usecase.Delete(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}
