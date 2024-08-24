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
	userAccRepo         repository.UserAccountRepo
	balanceMovementRepo repository.BalanceMovementRepo
}

func NewUserAccUC(
	userAccRepo repository.UserAccountRepo,
	balanceMovementRepo repository.BalanceMovementRepo,
) UserAccUC {
	return &userAccUC{
		userAccRepo:         userAccRepo,
		balanceMovementRepo: balanceMovementRepo,
	}
}

func (u *userAccUC) GetInfo(ctx context.Context, req *entity.UserAccGetInfoRequest) (*entity.UserAccGetInfoResponse, error) {
	account, err := u.userAccRepo.GetUserAccountByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	return &entity.UserAccGetInfoResponse{
		account.ToEntity(),
	}, nil
}

func (u *userAccUC) GetTransactionHistory(ctx context.Context, req *entity.UserAccGetTransactionHistoryRequest) (*entity.UserAccGetTransactionHistoryResponse, error) {
	_, err := u.userAccRepo.GetUserAccountByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	balanceMovementList, err := u.balanceMovementRepo.GetUserAccBalanceMovementList(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	var resp []*entity.BalanceMovement
	for _, bm := range balanceMovementList {
		resp = append(resp, bm.ToEntity())
	}

	return &entity.UserAccGetTransactionHistoryResponse{
		List: resp,
	}, nil
}
