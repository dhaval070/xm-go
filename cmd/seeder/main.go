package main

import (
	"log"
	"xmgo/pkg/config"
	"xmgo/pkg/models"
	"xmgo/pkg/repo"
	"xmgo/pkg/seeder"
)

func main() {
	cfg := config.LoadConfig()

	repository, err := repo.NewRepo(cfg.DSN)
	repository.AutoMigrate(models.Company{})

	if err != nil {
		log.Fatal(err)
	}

	seeder.Seed(repository)
}
