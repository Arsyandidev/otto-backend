package model

import "time"

type Voucher struct {
    ID           uint   `gorm:"primaryKey"`
    BrandID      uint   `gorm:"not null"`
    Code         string `gorm:"unique;not null"`
    Description  string `gorm:"size:255"`
    CostInPoints int    `gorm:"not null"`
    ValidFrom    string `gorm:"size:10"`
    ValidUntil   string `gorm:"size:10"`
    Brand        Brand  `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}