package model

import "time"

type Brand struct {
    ID          uint       `gorm:"primaryKey"`
    Name        string     `gorm:"not null"`
    Description string     `gorm:"size:255"`
    Vouchers    []Voucher  `gorm:"foreignKey:BrandID"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}