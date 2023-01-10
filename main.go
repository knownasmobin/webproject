package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	userHttp "git.ecobin.ir/ecomicro/template/app/user/delivery/http"
	userRepo "git.ecobin.ir/ecomicro/template/app/user/repository"
	userUsecase "git.ecobin.ir/ecomicro/template/app/user/usecase"
	"git.ecobin.ir/ecomicro/template/domain"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"

	"git.ecobin.ir/ecomicro/bootstrap/service"

	"git.ecobin.ir/ecomicro/bootstrap/config"
	conf "git.ecobin.ir/ecomicro/template/config"
	_ "git.ecobin.ir/ecomicro/template/docs"
	"git.ecobin.ir/ecomicro/transport"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type usecases struct {
	UserUsecase domain.UserUsecase
}

func setUsecase(db *gorm.DB, sf *sonyflake.Sonyflake) usecases {
	uiRepo := userRepo.NewUserRepository(db)
	uiUsecase := userUsecase.NewUserUsecase(uiRepo, sf)

	return usecases{
		UserUsecase: uiUsecase,
	}
}

func setHttpHandlers(g *gin.Engine, usecase usecases, httpConf transport.HTTPConfig, authMiddleware gin.HandlerFunc) {
	// handlers
	userHttp.NewUserHandler(g, authMiddleware, usecase.UserUsecase)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @title template Service API
// @version 1.0
// @description some desc
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	var myService *service.Service
	var c conf.Config
	var err error

	// Config file path
	wd, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	// Service setup
	myService, err = service.NewService("template", wd, config.JSON, &c)
	log.Fatal(err, "Failed to create new service!")

	dbConf := myService.Database["template"].GormDB
	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{StartTime: time.Now()})
	// Transport setup
	t, err := transport.NewTransport(myService)
	log.Fatal(err, "Failed to create new transport!")

	usecases := setUsecase(dbConf, sonyflake)
	httpConf := t.Config().Http["main"]
	// *****run http server*****
	_, err = t.HTTP("main", func(g *gin.Engine) {
		authMiddleware, err := t.AuthMiddleware()
		if err != nil {
			panic(err)
		}
		setHttpHandlers(g, usecases, httpConf, authMiddleware)
	})

	myService.Done()
}
