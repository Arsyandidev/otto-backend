package migration

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
	"fmt"
)

func MigrateCustomer() {
	if err := config.DB.AutoMigrate(&model.Voucher{}); err != nil {
		fmt.Println("Error Migrating Customers:", err)
	} else {
		fmt.Println("Customers Migrated Successfully")
	}
}