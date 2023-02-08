package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Usecase domain.Usecase
}

func NewUserHandler(g *gin.Engine, uu domain.Usecase) {
	rg := g.Group("/user")
	handler := &userHandler{
		Usecase: uu,
	}
	rg.GET("/",
		//  authMiddleware,
		handler.getUser)
	rg.PUT("/byCondition",
		//  authMiddleware,
		handler.getByCondition)
	rg.GET("/:id",
		//  authMiddleware,
		handler.getUserById)

	rg.POST("",
		//  authMiddleware,
		handler.createUser)
	rg.POST("/login",
		//  authMiddleware,
		handler.login)
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
	userId, err := strconv.Atoi(ctx.GetString("userId"))
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
// @Param id path int true "user id"
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

// login user godoc
// @Summary login user
// @Schemes
// @Description  login user
// @Tags user
// @Accept json
// @Produce json
// @Param body body LoginBody true "body params"
// @Success 200 {object} domain.User
// @Router /user/login [post]
func (uh *userHandler) login(ctx *gin.Context) {
	var body LoginBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	user, token, err := uh.Usecase.LoginUserCredential(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
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

// get by condition user godoc
// @Summary get by condition user
// @Schemes
// @Description  get by condition user
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdateUserBody true "body params"
// @Success 200 {array} domain.User
// @Router /user/byCondition [put]
func (uh *userHandler) getByCondition(ctx *gin.Context) {
	var body UpdateUserBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	user, err := uh.Usecase.GetByCondition(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
