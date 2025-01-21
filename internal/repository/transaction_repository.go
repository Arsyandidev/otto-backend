package repository

import (
	"be-otto-digital/internal/config"
	"be-otto-digital/internal/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
}

type transactionRepository struct {}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	return config.DB.Create(transaction).Error
}