package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/x"
	"git.ecobin.ir/services/template/domain"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Usecase domain.UserUsecase
}

func NewUserHandler(g *gin.Engine, authMiddleware gin.HandlerFunc, uu domain.UserUsecase) {
	rg := g.Group("/auth")
	handler := &userHandler{
		Usecase: uu,
	}
	rg.GET("/", authMiddleware, handler.getUser)
}

// get user godoc
// @Summary get user
// @Schemes
// @Description  get user
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {array} domain.User
// @Router /user [get]
func (uh *userHandler) getUser(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.GetString("userId"), 10, 64)
	if err != nil {
		return x.ErrHandler(ctx, err, errMap)
	}
	user, err := uh.Usecase.GetUserById(ctx, userId)
	if err != nil {
		x.ErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
