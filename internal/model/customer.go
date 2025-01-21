package model

import "time"

type Customer struct {
    ID           uint          `gorm:"primaryKey"`
    Name         string        `gorm:"not null"`
    Email        string        `gorm:"unique;not null"`
    Phone        string        `gorm:"size:15"`
    // Transactions []Transaction `gorm:"foreignKey:CustomerID"` // Relasi ke Transaction
    CreatedAt    time.Time     `gorm:"autoCreateTime"`
    UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
}