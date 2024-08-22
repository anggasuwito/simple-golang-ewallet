package usecase

import (
	"context"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/repository"
)

type TransactionUC interface {
	Topup(ctx context.Context, req *entity.TransactionTopupRequest) (*entity.TransactionTopupResponse, error)
	Transfer(ctx context.Context, req *entity.TransactionTransferRequest) (*entity.TransactionTransferResponse, error)
	Withdraw(ctx context.Context, req *entity.TransactionWithdrawRequest) (*entity.TransactionWithdrawResponse, error)
}

type transactionUC struct {
	transactionRepo repository.TransactionRepo
}

func NewTransactionUC(
	transactionRepo repository.TransactionRepo,
) TransactionUC {
	return &transactionUC{
		transactionRepo: transactionRepo,
	}
}

func (u *transactionUC) Topup(ctx context.Context, req *entity.TransactionTopupRequest) (*entity.TransactionTopupResponse, error) {
	return &entity.TransactionTopupResponse{}, nil
}

func (u *transactionUC) Transfer(ctx context.Context, req *entity.TransactionTransferRequest) (*entity.TransactionTransferResponse, error) {
	return &entity.TransactionTransferResponse{}, nil
}

func (u *transactionUC) Withdraw(ctx context.Context, req *entity.TransactionWithdrawRequest) (*entity.TransactionWithdrawResponse, error) {
	return &entity.TransactionWithdrawResponse{}, nil
}
