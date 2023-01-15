package grpc

import (
	"git.ecobin.ir/ecomicro/template/domain"
	"git.ecobin.ir/ecomicro/x"
	"google.golang.org/grpc/codes"
)

var errMap = x.ErrorMap[x.GrpcCustomError]{
	domain.ErrNotFound: {
		Status:  codes.NotFound,
		Message: "not found",
	},
}

// move blow func and types to x packages
