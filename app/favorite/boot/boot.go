package boot

import (
	"log"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	favoriteHttp "git.ecobin.ir/ecomicro/template/app/favorite/delivery/http"
	"git.ecobin.ir/ecomicro/template/app/favorite/domain"
	favoriteRepo "git.ecobin.ir/ecomicro/template/app/favorite/repository"
	favoriteUsecase "git.ecobin.ir/ecomicro/template/app/favorite/usecase"
	"git.ecobin.ir/ecomicro/transport"
	"git.ecobin.ir/ecomicro/x/structure"
	"gorm.io/gorm"
)

type favoriteBoot struct {
	transport *transport.Transport
	service   *service.Service
	db        *gorm.DB
}

var _ structure.BootInterface = &favoriteBoot{}

func (b *favoriteBoot) ApplyRepository(boot structure.Boot) {
	if _, ok := boot.Repositories[domain.DomainName]; ok {
		log.Fatalf("favorite repository already exist in repository map.")
	}
	boot.Repositories[domain.DomainName] = favoriteRepo.NewFavoriteRepository(b.db)
}

func (b *favoriteBoot) ApplyUsecase(boot structure.Boot) {
	favoriteRepository := structure.GetFromMap[domain.Repository](boot.Repositories, domain.DomainName)
	if _, ok := boot.Usecases[domain.DomainName]; ok {
		log.Fatalf("favorite usecase already exist in usecase map.")
	}
	boot.Usecases[domain.DomainName] = favoriteUsecase.NewFavoriteUsecase(favoriteRepository)
}
func (b *favoriteBoot) ApplyHttpHandler(boot structure.Boot) {
	favoriteUsecase := structure.GetFromMap[domain.Usecase](boot.Usecases, domain.DomainName)
	authMiddleware, err := b.transport.AuthMiddleware()
	if err != nil {
		log.Fatalf("fail to create auth middleware.")
	}
	favoriteHttp.NewFavoriteHandler(boot.Gin, authMiddleware, favoriteUsecase)
}
func (b *favoriteBoot) ApplyGrpcHandler(boot structure.Boot) {}

func (b *favoriteBoot) ApplyAdapters(boot structure.Boot) {
}

func Boot(service *service.Service, transport *transport.Transport) *favoriteBoot {
	return &favoriteBoot{
		transport: transport,
		service:   service,
		db:        service.Database["template"].GormDB,
	}
}
