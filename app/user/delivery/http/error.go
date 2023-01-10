package http

import (
	"net/http"

	"git.ecobin.ir/ecomicro/x"
	"git.ecobin.ir/services/template/domain"
)

var errMap = x.ErrorMap{
	domain.ErrUnprocessableEntity: {
		Status:       http.StatusUnprocessableEntity,
		CodeResponse: UnprocessableEntityResponse,
	},
}
