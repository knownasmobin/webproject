package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	userHttp "git.ecobin.ir/ecomicro/template/app/user/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/user/domain"
	userRepo "git.ecobin.ir/ecomicro/template/app/user/repository"
	userUsecase "git.ecobin.ir/ecomicro/template/app/user/usecase"
	"git.ecobin.ir/ecomicro/template/common"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"gorm.io/gorm"
)

type userBoot struct {
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
	usecase := userUsecase.NewUserUsecase(userRepository)
	boot.Usecases[domain.DomainName] = usecase
	common.NewAuth(usecase)
}
func (b *userBoot) ApplyHttpHandler(boot structure.Boot) {
	userUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	userHttp.NewUserHandler(boot.Gin, userUsecase)
}
func (b *userBoot) ApplyGrpcHandler(boot structure.Boot) {}
func (b *userBoot) ApplyAdapters(boot structure.Boot)    {}

func Boot(service *service.Service, transport *transport.Transport) *userBoot {
	return &userBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
