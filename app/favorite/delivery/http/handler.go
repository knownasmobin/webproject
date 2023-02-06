package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type favoriteHandler struct {
	Usecase domain.Usecase
}

func NewFavoriteHandler(g *gin.Engine, authMiddleware gin.HandlerFunc, uu domain.Usecase) {
	rg := g.Group("/favorite")
	handler := &favoriteHandler{
		Usecase: uu,
	}
	rg.GET("/",
		//  authMiddleware,
		handler.getFavorite)
	rg.GET("/:id",
		//  authMiddleware,
		handler.getFavoriteById)
	rg.POST("",
		//  authMiddleware,
		handler.createFavorite)
	rg.PUT("",
		//  authMiddleware,
		handler.updateFavorite)
}

// get favorite godoc
// @Summary get favorite
// @Schemes
// @Description  get favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {object} domain.Favorite
// @Router /favorite [get]
func (uh *favoriteHandler) getFavorite(ctx *gin.Context) {
	favoriteId, err := strconv.ParseUint(ctx.GetString("favoriteId"), 10, 64)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	favorite, err := uh.Usecase.GetFavoriteById(ctx, favoriteId)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

// get favorite by id godoc
// @Summary get favorite by id
// @Schemes
// @Description  get favorite by id
// @Tags favorite
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path uint64 true "favorite id"
// @Success 200 {object} domain.Favorite
// @Router /favorite/{id} [get]
func (uh *favoriteHandler) getFavoriteById(ctx *gin.Context) {

	var uri FavoriteIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	favorite, err := uh.Usecase.GetFavoriteById(ctx, uri.Id)
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
// @Param body body UpdateFavoriteBody true "body params"
// @Success 200 {object} domain.Favorite
// @Router /favorite [put]
func (uh *favoriteHandler) updateFavorite(ctx *gin.Context) {
	var body UpdateFavoriteBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	favorite, err := uh.Usecase.Update(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}
