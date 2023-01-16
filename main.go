package main

import (
	"log"
	"os"
	"path/filepath"

	"git.ecobin.ir/ecomicro/bootstrap/service"
	boot "git.ecobin.ir/ecomicro/template/boot"

	"git.ecobin.ir/ecomicro/bootstrap/config"
	conf "git.ecobin.ir/ecomicro/template/config"
	_ "git.ecobin.ir/ecomicro/template/docs"
)

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
	if err != nil {
		log.Fatal(err, "Failed to create new service!")
	}
	boot.Boot(myService)

	myService.Done()
}
