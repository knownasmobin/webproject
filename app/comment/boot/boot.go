package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	commentHttp "git.ecobin.ir/ecomicro/template/app/comment/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/comment/domain"
	commentRepo "git.ecobin.ir/ecomicro/template/app/comment/repository"
	commentUsecase "git.ecobin.ir/ecomicro/template/app/comment/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"gorm.io/gorm"
)

type commentBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &commentBoot{}

func (b *commentBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("comment repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = commentRepo.NewCommentRepository(b.db)
}

func (b *commentBoot) ApplyUsecase(boot structure.Boot) {
	commentRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("comment usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = commentUsecase.NewCommentUsecase(commentRepository)
}
func (b *commentBoot) ApplyHttpHandler(boot structure.Boot) {
	commentUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)

	commentHttp.NewCommentHandler(boot.Gin, commentUsecase)
}
func (b *commentBoot) ApplyGrpcHandler(boot structure.Boot) {}

func (b *commentBoot) ApplyAdapters(boot structure.Boot) {

}

func Boot(service *service.Service, transport *transport.Transport) *commentBoot {
	return &commentBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
