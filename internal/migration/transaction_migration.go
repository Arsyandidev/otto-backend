package migration

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
	"fmt"
)

func MigrateTransaction() {
	if err := config.DB.AutoMigrate(&model.Transaction{}); err != nil {
		fmt.Println("Error Migrating Transactions:", err)
	} else {
		fmt.Println("Transactions Migrated Successfully")
	}
}