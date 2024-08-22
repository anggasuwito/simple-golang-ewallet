package usecase

import (
	"context"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/repository"
)

type UserAccUC interface {
	GetInfo(ctx context.Context, req *entity.UserAccGetInfoRequest) (*entity.UserAccGetInfoResponse, error)
	GetTransactionHistory(ctx context.Context, req *entity.UserAccGetTransactionHistoryRequest) (*entity.UserAccGetTransactionHistoryResponse, error)
}

type userAccUC struct {
	userAccRepo repository.UserAccountRepo
}

func NewUserAccUC(
	userAccRepo repository.UserAccountRepo,
) UserAccUC {
	return &userAccUC{
		userAccRepo: userAccRepo,
	}
}

func (u *userAccUC) GetInfo(ctx context.Context, req *entity.UserAccGetInfoRequest) (*entity.UserAccGetInfoResponse, error) {
	return &entity.UserAccGetInfoResponse{}, nil
}

func (u *userAccUC) GetTransactionHistory(ctx context.Context, req *entity.UserAccGetTransactionHistoryRequest) (*entity.UserAccGetTransactionHistoryResponse, error) {
	return &entity.UserAccGetTransactionHistoryResponse{}, nil
}
