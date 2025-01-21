package service

import (
	"be-otto-digital/internal/model"
	"be-otto-digital/internal/repository"
	"be-otto-digital/internal/utils"
	"errors"
)

type TransactionService interface {
	CreateTransaction(transaction *model.Transaction) error
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s *transactionService) CreateTransaction(transaction *model.Transaction) error {
	if transaction.BrandID == 0 {
		errors.New("Brand ID Cannot be null")
	}

	if transaction.VoucherID == 0 {
		errors.New("Voucher ID Cannot be null")
	}

	if transaction.Customer.Email == "" || !utils.IsValidEmail(transaction.Customer.Email) {
		return errors.New("Valid Email is required")
	}

 	return s.transactionRepo.Create(transaction)
}
