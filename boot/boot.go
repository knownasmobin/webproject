package boot

import (
	"log"
	"time"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	baz "git.ecobin.ir/ecomicro/template/app/baz/boot"
	user "git.ecobin.ir/ecomicro/template/app/user/boot"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func Boot(service *service.Service) {

	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{StartTime: time.Now()})

	t, err := transport.NewTransport(service)
	if err != nil {
		log.Fatal(err, "Failed to create new transport!")
	}

	boots := make([]structure.BootInterface, 0)
	// boot domains
	boots = append(boots, baz.Boot(service, t))
	boots = append(boots, user.Boot(service, sonyflake, t))

	bootData := structure.Boot{
		GrpcServers:  make(map[string]interface{}),
		GrpcClients:  make(map[string]interface{}),
		Usecases:     make(map[string]interface{}),
		Repositories: make(map[string]interface{}),
		Adapters:     make(map[string]interface{}),
	}
	bootData.GrpcClients["foo"], err = t.GRPCClient("foo")
	_, err = t.HTTP("main", func(g *gin.Engine) { bootData.Gin = g })

	// repository
	for _, boot := range boots {
		boot.ApplyRepository(bootData)
	}
	// usecase
	for _, boot := range boots {
		boot.ApplyUsecase(bootData)
	}
	// http
	for _, boot := range boots {
		boot.ApplyHttpHandler(bootData)
	}
	// grpc
	_, err = t.GRPCSevrer("user", func(g *grpc.Server) {
		bootData.GrpcServers["user"] = g
		for _, boot := range boots {
			boot.ApplyGrpcHandler(bootData)
		}
	})
	// adapter
	for _, boot := range boots {
		boot.ApplyAdapters(bootData)
	}
	// swagger
	bootData.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
