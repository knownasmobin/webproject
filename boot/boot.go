package boot

import (
	"log"
	"time"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	bazRepo "git.ecobin.ir/ecomicro/template/app/baz/repository"
	bazUsecase "git.ecobin.ir/ecomicro/template/app/baz/usecase"
	userGRPC "git.ecobin.ir/ecomicro/template/app/user/delivery/grpc"
	userHttp "git.ecobin.ir/ecomicro/template/app/user/delivery/http"
	userRepo "git.ecobin.ir/ecomicro/template/app/user/repository"
	userUsecase "git.ecobin.ir/ecomicro/template/app/user/usecase"
	"git.ecobin.ir/ecomicro/template/domain"
	userBazAdapter "git.ecobin.ir/ecomicro/template/infra/adapters/baz"
	userFooAdapter "git.ecobin.ir/ecomicro/template/infra/adapters/foo"
	"git.ecobin.ir/ecomicro/transport"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type usecases struct {
	UserUsecase domain.UserUsecase
	BazUsecase  domain.BazUsecase
}

func setUsecase(db *gorm.DB, sf *sonyflake.Sonyflake, fooConnection *grpc.ClientConn) usecases {
	userRepo := userRepo.NewUserRepository(db)
	bazRepo := bazRepo.NewBazRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepo, sf)
	bazUsecase := bazUsecase.NewBazUsecase(bazRepo)
	fooAdapter := userFooAdapter.NewFooAdapter(fooConnection)
	bazAdapter := userBazAdapter.NewBazUsecaseAdapter(bazUsecase)

	userUsecase.SetAdapters(fooAdapter, bazAdapter)
	return usecases{
		UserUsecase: userUsecase,
		BazUsecase:  bazUsecase,
	}
}

func setHttpHandlers(g *gin.Engine, usecase usecases, httpConf transport.HTTPConfig, authMiddleware gin.HandlerFunc) {
	// handlers
	userHttp.NewUserHandler(g, authMiddleware, usecase.UserUsecase)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setGRPCHandlers(s *grpc.Server, usecase usecases) {
	// handlers
	userGRPC.NewUserHandler(s, usecase.UserUsecase)
}

func Boot(service *service.Service) {

	dbConf := service.Database["template"].GormDB
	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{StartTime: time.Now()})
	// Transport setup
	t, err := transport.NewTransport(service)
	if err != nil {
		log.Fatal(err, "Failed to create new transport!")
	}
	fooGrpcConnection, err := t.GRPCClient("foo")
	if err != nil {
		log.Fatal("fail to dial grpc client : ", err)
	}
	usecases := setUsecase(dbConf, sonyflake, fooGrpcConnection)
	httpConf := t.Config().Http["main"]
	// *****run http server*****
	_, err = t.HTTP("main", func(g *gin.Engine) {
		authMiddleware, err := t.AuthMiddleware()
		if err != nil {
			log.Fatal("fail to run http server : ", err)
		}
		setHttpHandlers(g, usecases, httpConf, authMiddleware)
	})
	// ***** run grpc server ****
	_, err = t.GRPCSevrer("template", func(g *grpc.Server) {
		setGRPCHandlers(g, usecases)
	})

}
