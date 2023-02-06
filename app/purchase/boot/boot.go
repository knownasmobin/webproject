package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	purchaseHttp "git.ecobin.ir/ecomicro/template/app/purchase/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/purchase/domain"
	purchaseRepo "git.ecobin.ir/ecomicro/template/app/purchase/repository"
	purchaseUsecase "git.ecobin.ir/ecomicro/template/app/purchase/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
)

type purchaseBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &purchaseBoot{}

func (b *purchaseBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("purchase repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = purchaseRepo.NewPurchaseRepository(b.db)
}

func (b *purchaseBoot) ApplyUsecase(boot structure.Boot) {
	purchaseRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("purchase usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = purchaseUsecase.NewPurchaseUsecase(purchaseRepository)
}
func (b *purchaseBoot) ApplyHttpHandler(boot structure.Boot) {
	purchaseUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	authMiddleware, err := b.transport.AuthMiddleware()
	if err != nil {
		log.Fatalf("fail to create auth middleware.")
	}
	purchaseHttp.NewPurchaseHandler(boot.Gin, authMiddleware, purchaseUsecase)
}
func (b *purchaseBoot) ApplyGrpcHandler(boot structure.Boot) {
}

func (b *purchaseBoot) ApplyAdapters(boot structure.Boot) {

}

func Boot(service *service.Service, sonyflake *sonyflake.Sonyflake, transport *transport.Transport) *purchaseBoot {
	return &purchaseBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
