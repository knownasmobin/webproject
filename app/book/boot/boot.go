package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	bookHttp "git.ecobin.ir/ecomicro/template/app/book/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/book/domain"
	bookRepo "git.ecobin.ir/ecomicro/template/app/book/repository"
	bookUsecase "git.ecobin.ir/ecomicro/template/app/book/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"gorm.io/gorm"
)

type bookBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &bookBoot{}

func (b *bookBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("book repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = bookRepo.NewBookRepository(b.db)
}

func (b *bookBoot) ApplyUsecase(boot structure.Boot) {
	bookRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("book usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = bookUsecase.NewBookUsecase(bookRepository)
}
func (b *bookBoot) ApplyHttpHandler(boot structure.Boot) {
	bookUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	authMiddleware, err := b.transport.AuthMiddleware()
	if err != nil {
		log.Fatalf("fail to create auth middleware.")
	}
	bookHttp.NewBookHandler(boot.Gin, authMiddleware, bookUsecase)
}
func (b *bookBoot) ApplyGrpcHandler(boot structure.Boot) {}

func (b *bookBoot) ApplyAdapters(boot structure.Boot) {}

func Boot(service *service.Service, transport *transport.Transport) *bookBoot {
	return &bookBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
