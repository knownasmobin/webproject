package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type purchaseHandler struct {
	Usecase domain.Usecase
}

func NewPurchaseHandler(g *gin.Engine, uu domain.Usecase) {
	rg := g.Group("/purchase")
	handler := &purchaseHandler{
		Usecase: uu,
	}
	rg.POST("",
		//  authMiddleware,
		handler.createPurchase)
	rg.PUT("",
		//  authMiddleware,
		handler.getByCondition)
}

// create purchase godoc
// @Summary create purchase
// @Schemes
// @Description  create purchase
// @Tags purchase
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body CreatePurchaseBody true "body params"
// @Success 200 {object} domain.Purchase
// @Router /purchase [post]
func (uh *purchaseHandler) createPurchase(ctx *gin.Context) {
	var body CreatePurchaseBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	purchase, err := uh.Usecase.Create(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, purchase)
}

// update purchase godoc
// @Summary update purchase
// @Schemes
// @Description  update purchase
// @Tags purchase
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param body body UpdatePurchaseBody true "body params"
// @Success 200 {array} domain.Purchase
// @Router /purchase [put]
func (uh *purchaseHandler) getByCondition(ctx *gin.Context) {
	var body UpdatePurchaseBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	purchase, err := uh.Usecase.GetByCondition(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, purchase)
}
