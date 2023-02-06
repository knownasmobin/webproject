package http

import (
	"net/http"
	"strconv"

	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
	"git.ecobin.ir/ecomicro/x"

	"github.com/gin-gonic/gin"
)

type purchaseHandler struct {
	Usecase domain.Usecase
}

func NewPurchaseHandler(g *gin.Engine, authMiddleware gin.HandlerFunc, uu domain.Usecase) {
	rg := g.Group("/purchase")
	handler := &purchaseHandler{
		Usecase: uu,
	}
	rg.GET("/",
		//  authMiddleware,
		handler.getPurchase)
	rg.GET("/:id",
		//  authMiddleware,
		handler.getPurchaseById)
	rg.POST("",
		//  authMiddleware,
		handler.createPurchase)
	rg.PUT("",
		//  authMiddleware,
		handler.updatePurchase)
}

// get purchase godoc
// @Summary get purchase
// @Schemes
// @Description  get purchase
// @Tags purchase
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Success 200 {object} domain.Purchase
// @Router /purchase [get]
func (uh *purchaseHandler) getPurchase(ctx *gin.Context) {
	purchaseId, err := strconv.ParseUint(ctx.GetString("purchaseId"), 10, 64)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	purchase, err := uh.Usecase.GetPurchaseById(ctx, purchaseId)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, purchase)
}

// get purchase by id godoc
// @Summary get purchase by id
// @Schemes
// @Description  get purchase by id
// @Tags purchase
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer jwtToken"
// @Param id path uint64 true "purchase id"
// @Success 200 {object} domain.Purchase
// @Router /purchase/{id} [get]
func (uh *purchaseHandler) getPurchaseById(ctx *gin.Context) {

	var uri PurchaseIdUri
	err := ctx.BindUri(&uri)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}

	purchase, err := uh.Usecase.GetPurchaseById(ctx, uri.Id)
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, purchase)
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
// @Success 200 {object} domain.Purchase
// @Router /purchase [put]
func (uh *purchaseHandler) updatePurchase(ctx *gin.Context) {
	var body UpdatePurchaseBody
	err := ctx.Bind(&body)
	if err != nil {
		x.HttpErrHandler(ctx, domain.ErrUnprocessableEntity, errMap)
		return
	}
	purchase, err := uh.Usecase.Update(ctx, body.toDomain())
	if err != nil {
		x.HttpErrHandler(ctx, err, errMap)
		return
	}
	ctx.JSON(http.StatusOK, purchase)
}
