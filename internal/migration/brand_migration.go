package migration

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
	"fmt"
)

func MigrateBrand() {
	if err := config.DB.AutoMigrate(&model.Brand{}); err != nil {
		fmt.Println("Error Migrating Brands:", err)
	} else {
		fmt.Println("Brands Migrated Successfully")
	}
}