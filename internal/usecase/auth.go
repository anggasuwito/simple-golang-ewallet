package usecase

import (
	"context"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/repository"
)

type AuthUC interface {
	LoginPIN(ctx context.Context, req *entity.AuthLoginPINRequest) (*entity.AuthLoginPINResponse, error)
	VerifyPIN(ctx context.Context, req *entity.AuthVerifyPINRequest) (*entity.AuthVerifyPINResponse, error)
}

type authUC struct {
	userRepo repository.UserRepo
}

func NewAuthUC(
	userRepo repository.UserRepo,
) AuthUC {
	return &authUC{
		userRepo: userRepo,
	}
}

func (u *authUC) LoginPIN(ctx context.Context, req *entity.AuthLoginPINRequest) (*entity.AuthLoginPINResponse, error) {
	return &entity.AuthLoginPINResponse{}, nil
}

func (u *authUC) VerifyPIN(ctx context.Context, req *entity.AuthVerifyPINRequest) (*entity.AuthVerifyPINResponse, error) {
	return &entity.AuthVerifyPINResponse{}, nil
}
