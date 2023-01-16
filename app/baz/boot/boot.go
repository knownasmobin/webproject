package boot

import (
	"fmt"
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	"git.ecobin.ir/ecomicro/template/app/baz/domain"
	bazRepo "git.ecobin.ir/ecomicro/template/app/baz/repository"
	bazUsecase "git.ecobin.ir/ecomicro/template/app/baz/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
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

type bazBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &bazBoot{}

func (b *bazBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("baz repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = bazRepo.NewBazRepository(b.db)
}
func (b *bazBoot) ApplyUsecase(boot structure.Boot) {
	bazRepository := getFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("baz usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = bazUsecase.NewBazUsecase(bazRepository)
}
func (b *bazBoot) ApplyHttpHandler(boot structure.Boot) {}
func (b *bazBoot) ApplyGrpcHandler(boot structure.Boot) {}
func (b *bazBoot) ApplyAdapters(boot structure.Boot)    {}

func Boot(service *service.Service, transport *transport.Transport) *bazBoot {
	return &bazBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
