package main

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/handler"
	"be-otto-digital/internal/migration"
	"be-otto-digital/internal/repository"
	routes "be-otto-digital/internal/router"
	"be-otto-digital/internal/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	migration.MigrateBrand()
	migration.MigrateVoucher()
	migration.MigrateCustomer()
	migration.MigrateTransaction()

	brandRepo := repository.NewBrandRepository()
	voucherRepo := repository.NewVoucherRepository()
	transactionRepo := repository.NewTransactionRepository()
	

	brandService := service.NewBrandService(brandRepo)
	voucherService := service.NewVoucherService(voucherRepo)
	transactionService := service.NewTransactionService(transactionRepo)
	
	brandHandler := handler.NewBrandHandler(brandService)
	voucherHandler := handler.NewVoucherHandler(voucherService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	
	r := mux.NewRouter()
	routes.RegisterRoutes(r, brandHandler, voucherHandler, transactionHandler)

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", r)
}