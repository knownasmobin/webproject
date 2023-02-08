package common

import (
	"context"
	"net/http"
	"strings"

	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"github.com/gin-gonic/gin"
)

type GuardAdapter interface {
	ValidateToken(ctx context.Context, tokenStr string) (*domain.User, error)
}

type Auth struct {
	g GuardAdapter
}

var auth *Auth

func NewAuth(g GuardAdapter) {
	auth = &Auth{
		g: g,
	}
}
func GetAuth() *Auth {
	return auth
}
func (a *Auth) Guard(justAdmin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := tooty.OpenAnAPMSpan(ctx.Request.Context(), "[M] Auth", "middleware")
		defer tooty.CloseTheAPMSpan(span)
		var authHeader string
		if header := ctx.Request.Header["authorization"]; len(header) > 0 {
			authHeader = header[0]
		}
		if authHeader == "" {
			if header := ctx.Request.Header["Authorization"]; len(header) > 0 {
				authHeader = header[0]
			}
		}

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			ctx.Abort()
			return
		}

		splitted := strings.Split(authHeader, " ")
		if len(splitted) < 2 || splitted[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			ctx.Abort()
			return
		}
		accessToken := splitted[1]

		user, err := a.g.ValidateToken(ctx, accessToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			ctx.Abort()
			return
		}
		if justAdmin && !user.IsAdmin {
			ctx.JSON(http.StatusForbidden, gin.H{})
			ctx.Abort()
			return
		}
		ctx.Set("isAdmin", user.IsAdmin)
		ctx.Set("userId", user.Id)
		ctx.Set("token", accessToken)
		tooty.SetTransactionLabel(ctx.Request.Context(), "user.id", user.Id)
		tooty.CloseTheAPMSpan(span)
		ctx.Next()
	}
}
