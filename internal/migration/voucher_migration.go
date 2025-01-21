package migration

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
	"fmt"
)

func MigrateVoucher() {
	if err := config.DB.AutoMigrate(&model.Voucher{}); err != nil {
		fmt.Println("Error Migrating Vouchers:", err)
	} else {
		fmt.Println("Vouchers Migrated Successfully")
	}
}