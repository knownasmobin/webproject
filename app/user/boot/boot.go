package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	bazDomain "git.ecobin.ir/ecomicro/template/app/baz/domain"
	userGRPC "git.ecobin.ir/ecomicro/template/app/user/delivery/grpc"
	userHttp "git.ecobin.ir/ecomicro/template/app/user/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/user/domain"
	userRepo "git.ecobin.ir/ecomicro/template/app/user/repository"
	userUsecase "git.ecobin.ir/ecomicro/template/app/user/usecase"
	userBazAdapter "git.ecobin.ir/ecomicro/template/infra/adapters/baz"
	userFooAdapter "git.ecobin.ir/ecomicro/template/infra/adapters/foo"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type userBoot struct {
	sonyflake *sonyflake.Sonyflake
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &userBoot{}

func (b *userBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("user repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = userRepo.NewRepository(b.db)
}

func (b *userBoot) ApplyUsecase(boot structure.Boot) {
	userRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("user usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = userUsecase.NewUserUsecase(userRepository, b.sonyflake)
}
func (b *userBoot) ApplyHttpHandler(boot structure.Boot) {
	userUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	authMiddleware, err := b.transport.AuthMiddleware()
	if err != nil {
		log.Fatalf("fail to create auth middleware.")
	}
	userHttp.NewUserHandler(boot.Gin, authMiddleware, userUsecase)
}
func (b *userBoot) ApplyGrpcHandler(boot structure.Boot) {
	userUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	grpcServer := structure.GetFromMap[*grpc.Server](boot.GrpcServers, domain.DomainName)
	userGRPC.NewUserHandler(grpcServer, userUsecase)
}
func (b *userBoot) ApplyAdapters(boot structure.Boot) {
	userUsecase := structure.GetFromMap[domain.Adapter](boot.Usecases, domain.DomainName)
	bazUsecase := structure.GetFromMap[bazDomain.Usecase](boot.Usecases, bazDomain.DomainName)
	fooGrpcClient := structure.GetFromMap[*grpc.ClientConn](boot.GrpcClients, "foo")

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
