package repository

import "gorm.io/gorm"

type TransactionRepo interface {
}

type transactionRepo struct {
	masterDB *gorm.DB
}

func NewTransactionRepo(masterDB *gorm.DB) TransactionRepo {
	return &transactionRepo{
		masterDB: masterDB,
	}
}
