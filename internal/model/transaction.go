package model

import "time"

type Transaction struct {
    ID         uint `gorm:"primaryKey"`
    BrandID    uint
    Brand      Brand `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    VoucherID  uint
    Voucher    Voucher `gorm:"foreignKey:VoucherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Customer   Customer `gorm:"embedded"`
    TotalPoint uint
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}