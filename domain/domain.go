package domain

import (
	"github.com/gin-gonic/gin"
)

type DomainBoot struct {
	GrpcServers  map[string]interface{}
	GrpcClients  map[string]interface{}
	Gin          *gin.Engine
	Usecases     map[string]interface{}
	Repositories map[string]interface{}
	Adapters     map[string]interface{}
}
type Boot interface {
	ApplyRepository(boot DomainBoot)
	ApplyUsecase(boot DomainBoot)
	ApplyHttpHandler(boot DomainBoot)
	ApplyGrpcHandler(boot DomainBoot)
	ApplyAdapters(boot DomainBoot)
}
