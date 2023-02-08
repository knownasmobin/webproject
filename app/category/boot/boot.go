package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	categoryHttp "git.ecobin.ir/ecomicro/template/app/category/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/category/domain"
	categoryRepo "git.ecobin.ir/ecomicro/template/app/category/repository"
	categoryUsecase "git.ecobin.ir/ecomicro/template/app/category/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"gorm.io/gorm"
)

type categoryBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &categoryBoot{}

func (b *categoryBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("category repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = categoryRepo.NewCategoryRepository(b.db)
}

func (b *categoryBoot) ApplyUsecase(boot structure.Boot) {
	categoryRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("category usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = categoryUsecase.NewCategoryUsecase(categoryRepository)
}
func (b *categoryBoot) ApplyHttpHandler(boot structure.Boot) {
	categoryUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	categoryHttp.NewCategoryHandler(boot.Gin, categoryUsecase)
}
func (b *categoryBoot) ApplyGrpcHandler(boot structure.Boot) {
}

func (b *categoryBoot) ApplyAdapters(boot structure.Boot) {}
func (b *categoryBoot) ApplyOthers(boot structure.Boot)   {}
func (b *categoryBoot) ApplyCronjobs(boot structure.Boot) {}

func Boot(service *service.Service, transport *transport.Transport) *categoryBoot {
	return &categoryBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
