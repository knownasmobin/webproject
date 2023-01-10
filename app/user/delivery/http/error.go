package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/x"
)

var errMap = x.ErrorMap{
	domain.ErrUnprocessableEntity: {
		Status:       http.StatusUnprocessableEntity,
		CodeResponse: UnprocessableEntityResponse,
	},
}
