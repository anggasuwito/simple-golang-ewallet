package usecase

import (
	"context"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/repository"
)

type UserUC interface {
	GetInfo(ctx context.Context, req *entity.UserGetInfoRequest) (*entity.UserGetInfoResponse, error)
	GetTransactionHistory(ctx context.Context, req *entity.UserGetTransactionHistoryRequest) (*entity.UserGetTransactionHistoryResponse, error)
}

type userUC struct {
	userRepo repository.UserRepo
}

func NewUserUC(
	userRepo repository.UserRepo,
) UserUC {
	return &userUC{
		userRepo: userRepo,
	}
}

func (u *userUC) GetInfo(ctx context.Context, req *entity.UserGetInfoRequest) (*entity.UserGetInfoResponse, error) {
	return &entity.UserGetInfoResponse{}, nil
}

func (u *userUC) GetTransactionHistory(ctx context.Context, req *entity.UserGetTransactionHistoryRequest) (*entity.UserGetTransactionHistoryResponse, error) {
	return &entity.UserGetTransactionHistoryResponse{}, nil
}
