package main

import (
	"xmgo/api/http/companies"
	usecase "xmgo/pkg/usecase/companies"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	uc := usecase.NewUseCase()
	companiesHandler := companies.NewHandler(uc)
	companiesHandler.InitRoutes(router.Group("/companies"))

	router.Run(":8080")
}
