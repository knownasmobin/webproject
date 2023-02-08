package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/template/app/comment/domain"
	"git.ecobin.ir/ecomicro/x"
)

var (
	UnprocessableEntityResponse = x.CodeResponse{
		Code:    "unprocessable-entity",
		Message: "unprocessable entity.",
	}

	errMap = x.ErrorMap[x.HttpCustomError]{
		domain.ErrUnprocessableEntity: {
			Status:       http.StatusUnprocessableEntity,
			CodeResponse: UnprocessableEntityResponse,
		},
	}
)
