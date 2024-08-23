package usecase

import (
	"context"
	"simple-golang-ewallet/internal/constant"
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
	pinRepo     repository.PINRepo
}

func NewAuthUC(
	accountRepo repository.UserAccountRepo,
	pinRepo repository.PINRepo,
) AuthUC {
	return &authUC{
		accountRepo: accountRepo,
		pinRepo:     pinRepo,
	}
}

func (u *authUC) LoginPIN(ctx context.Context, req *entity.AuthLoginPINRequest) (*entity.AuthLoginPINResponse, error) {
	account, err := u.accountRepo.GetUserAccountByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	validPin := utils.CompareHashCredential(req.PIN, account.PIN)
	if !validPin {
		return nil, utils.ErrBadRequest("Invalid pin", "authUC.LoginPIN.CompareHashCredential")
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
	pinType := []string{constant.PINTypeTransfer, constant.PINTypeWithdraw}
	if !utils.InArray(req.Type, pinType) {
		return nil, utils.ErrBadRequest("Invalid type", "authUC.VerifyPIN.Type")
	}

	account, err := u.accountRepo.GetUserAccountByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	validPin := utils.CompareHashCredential(req.PIN, account.PIN)
	if !validPin {
		return nil, utils.ErrBadRequest("Invalid pin", "authUC.VerifyPIN.CompareHashCredential")
	}

	err = u.pinRepo.SetVerifiedPINByTypeCache(ctx, account.ID, req.Type)
	if err != nil {
		return nil, err
	}
	return &entity.AuthVerifyPINResponse{}, nil
}
