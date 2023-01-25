package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	"git.ecobin.ir/ecomicro/bootstrap/store/redis"
	"git.ecobin.ir/ecomicro/template/app/baz/domain"
	bazRepo "git.ecobin.ir/ecomicro/template/app/baz/repository/psql"
	redisRepo "git.ecobin.ir/ecomicro/template/app/baz/repository/redis"
	bazUsecase "git.ecobin.ir/ecomicro/template/app/baz/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"gorm.io/gorm"
)

type bazBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
	rdb       *redis.Redis
}

var _ structure.BootInterface = &bazBoot{}

func (b *bazBoot) ApplyRepository(boot structure.Boot) {

	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("baz repository already exist in repository map.")
	}
	psqlRepository := bazRepo.NewRepository(b.db)
	// create redis cache layer
	redisCacheRepository := redisRepo.New(psqlRepository, b.rdb)
	boot.Repositories[domain.DomainName] = redisCacheRepository

}
func (b *bazBoot) ApplyUsecase(boot structure.Boot) {
	bazRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("baz usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = bazUsecase.NewBazUsecase(bazRepository)
}
func (b *bazBoot) ApplyHttpHandler(boot structure.Boot) {}
func (b *bazBoot) ApplyGrpcHandler(boot structure.Boot) {}
func (b *bazBoot) ApplyAdapters(boot structure.Boot)    {}

func Boot(service *service.Service, transport *transport.Transport) *bazBoot {
	rdb, err := service.Store.Redis("main")
	if err != nil {
		log.Fatalf("baz domain need redis: %+v", err)
	}
	return &bazBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
		rdb:       rdb,
	}
}
