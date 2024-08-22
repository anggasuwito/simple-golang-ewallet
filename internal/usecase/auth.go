package usecase

import (
	"context"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/repository"
	"simple-golang-ewallet/internal/utils"
)

type AuthUC interface {
	LoginPIN(ctx context.Context, req *entity.AuthLoginPINRequest) (*entity.AuthLoginPINResponse, error)
	VerifyPIN(ctx context.Context, req *entity.AuthVerifyPINRequest) (*entity.AuthVerifyPINResponse, error)
}

type authUC struct {
	accountRepo repository.UserAccountRepo
}

func NewAuthUC(
	accountRepo repository.UserAccountRepo,
) AuthUC {
	return &authUC{
		accountRepo: accountRepo,
	}
}

func (u *authUC) LoginPIN(ctx context.Context, req *entity.AuthLoginPINRequest) (*entity.AuthLoginPINResponse, error) {
	account, err := u.accountRepo.GetUserAccount(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	token, _, err := utils.GenerateJWT(account)
	if err != nil {
		return nil, utils.ErrInternal("Failed generate jwt : "+err.Error(), "authUC.LoginPIN.GenerateJWT")
	}

	return &entity.AuthLoginPINResponse{
		AccessToken: token,
	}, nil
}

func (u *authUC) VerifyPIN(ctx context.Context, req *entity.AuthVerifyPINRequest) (*entity.AuthVerifyPINResponse, error) {
	return &entity.AuthVerifyPINResponse{}, nil
}
