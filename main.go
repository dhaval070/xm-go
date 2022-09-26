package main

import (
	"log"
	"xmgo/api/handlers"
	"xmgo/pkg/config"
	"xmgo/pkg/models"
	"xmgo/pkg/repo"
	"xmgo/pkg/usecase/companies"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cfg := config.LoadConfig()
	log.Println(cfg)

	repository, err := repo.NewRepo(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	repository.AutoMigrate(models.Company{})

	uc := companies.NewUseCase(repository)
	handler := handlers.NewHandler(uc)
	handler.InitRoutes(router.Group("/companies"))

	router.Run(":8080")
}
