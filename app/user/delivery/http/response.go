package http

import "git.ecobin.ir/ecomicro/x"

var (
	UnprocessableEntityResponse = x.CodeResponse{
		Code:    "unprocessable-entity",
		Message: "unprocessable entity.",
	}
)
