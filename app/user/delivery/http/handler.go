package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Usecase domain.UserUsecase
}

func NewUserHandler(g *gin.Engine, authMiddleware gin.HandlerFunc, uu domain.UserUsecase) {
	rg := g.Group("/user")
	handler := &userHandler{
		Usecase: uu,
	}
	rg.GET("/",
		//  authMiddleware,
		handler.getUser)
	rg.GET("/:id",
		//  authMiddleware,
		handler.getUserById)
	rg.POST("",
		//  authMiddleware,
		handler.createUser)
	rg.PUT("",
		//  authMiddleware,
		handler.updateUser)
}

// get user godoc
// @Summary get user
// @Schemes
// @Description  get user
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {object} domain.User
// @Router /user [get]
func (uh *userHandler) getUser(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.GetString("userId"), 10, 64)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	user, err := uh.Usecase.GetUserById(ctx, userId)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// get user by id godoc
// @Summary get user by id
// @Schemes
// @Description  get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path uint64 true "user id"
// @Success 200 {object} domain.User
// @Router /user/{id} [get]
func (uh *userHandler) getUserById(ctx *gin.Context) {

	var uri UserIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	user, err := uh.Usecase.GetUserById(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// create user godoc
// @Summary create user
// @Schemes
// @Description  create user
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body CreateUserBody true "body params"
// @Success 200 {object} domain.User
// @Router /user [post]
func (uh *userHandler) createUser(ctx *gin.Context) {
	var body CreateUserBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	user, err := uh.Usecase.Create(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// update user godoc
// @Summary update user
// @Schemes
// @Description  update user
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdateUserBody true "body params"
// @Success 200 {object} domain.User
// @Router /user [put]
func (uh *userHandler) updateUser(ctx *gin.Context) {
	var body UpdateUserBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	user, err := uh.Usecase.Update(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
