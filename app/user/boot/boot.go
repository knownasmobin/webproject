package boot

import (
	"fmt"
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	bazDomain "git.ecobin.ir/ecomicro/template/app/baz/domain"
	userGRPC "git.ecobin.ir/ecomicro/template/app/user/delivery/grpc"
	userHttp "git.ecobin.ir/ecomicro/template/app/user/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/user/domain"
	userRepo "git.ecobin.ir/ecomicro/template/app/user/repository"
	userUsecase "git.ecobin.ir/ecomicro/template/app/user/usecase"
	mainDomain "git.ecobin.ir/ecomicro/template/domain"
	userBazAdapter "git.ecobin.ir/ecomicro/template/infra/adapters/baz"
	userFooAdapter "git.ecobin.ir/ecomicro/template/infra/adapters/foo"
	"git.ecobin.ir/ecomicro/transport"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func getFromMap[T any](stringMap map[string]interface{}, key string) T {
	if value, ok := stringMap[key]; ok {
		if v, ok := value.(T); ok {
			return v
		}
		panic(fmt.Sprintf("assertion failed: %+v", stringMap[key]))
	}
	panic(fmt.Sprintf("key not found: map is=> %+v  and key is %+v", stringMap, key))
}

type userBoot struct {
	sonyflake *sonyflake.Sonyflake
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ mainDomain.Boot = &userBoot{}

func (b *userBoot) ApplyRepository(boot mainDomain.DomainBoot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("user repository already exist in repository map.")
	}
	boot.Repositories["user"] = userRepo.NewUserRepository(b.db)
}

func (b *userBoot) ApplyUsecase(boot mainDomain.DomainBoot) {
	userRepository := getFromMap[domain.Repository](boot.Repositories, "user")
	if _, ok := boot.Usecases["user"]; ok {
		log.Fatalf("user usecase already exist in usecase map.")
	}
	boot.Usecases["user"] = userUsecase.NewUserUsecase(userRepository, b.sonyflake)
}
func (b *userBoot) ApplyHttpHandler(boot mainDomain.DomainBoot) {
	userUsecase := getFromMap[domain.Usecase](boot.Usecases, "user")
	authMiddleware, err := b.transport.AuthMiddleware()
	if err != nil {
		log.Fatalf("fail to create auth middleware.")
	}
	userHttp.NewUserHandler(boot.Gin, authMiddleware, userUsecase)
}
func (b *userBoot) ApplyGrpcHandler(boot mainDomain.DomainBoot) {
	userUsecase := getFromMap[domain.Usecase](boot.Usecases, "user")
	grpcServer := getFromMap[*grpc.Server](boot.GrpcServers, "user")
	userGRPC.NewUserHandler(grpcServer, userUsecase)
}
func (b *userBoot) ApplyAdapters(boot mainDomain.DomainBoot) {
	userUsecase := getFromMap[domain.Adapter](boot.Usecases, "user")
	bazUsecase := getFromMap[bazDomain.Usecase](boot.Usecases, "baz")
	fooGrpcClient := getFromMap[*grpc.ClientConn](boot.GrpcClients, "foo")

	fooAdapter := userFooAdapter.NewFooAdapter(fooGrpcClient)
	bazAdapter := userBazAdapter.NewBazUsecaseAdapter(bazUsecase)

	userUsecase.SetAdapters(fooAdapter, bazAdapter)
}

func Boot(service *service.Service, sonyflake *sonyflake.Sonyflake, transport *transport.Transport) *userBoot {
	return &userBoot{
		sonyflake: sonyflake,
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
