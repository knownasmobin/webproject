package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/x"
)

var errMap = x.ErrorMap[x.HttpCustomError]{
	domain.ErrUnprocessableEntity: {
		Status:       http.StatusUnprocessableEntity,
		CodeResponse: UnprocessableEntityResponse,
	},
}
