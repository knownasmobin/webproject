package grpc

import (
	"git.ecobin.ir/ecomicro/template/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errMap = GrpcErrorMap{
	domain.ErrNotFound: {
		Code:    codes.NotFound,
		Message: "not found",
	},
}

// move blow func and types to x packages
type GrpcCustomError struct {
	Code    codes.Code
	Message string
}

type GrpcErrorMap map[error]GrpcCustomError

func GrpcResponseError(err error, errMap GrpcErrorMap) error {
	if v, ok := errMap[err]; ok {
		return status.Error(v.Code, v.Message)
	} else {
		return status.Error(codes.Internal, "internal server error")
	}
}
