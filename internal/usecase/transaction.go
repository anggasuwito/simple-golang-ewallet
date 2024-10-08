package usecase

import (
	"context"
	"simple-golang-ewallet/internal/constant"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/domain/model"
	"simple-golang-ewallet/internal/repository"
	"simple-golang-ewallet/internal/utils"
)

type TransactionUC interface {
	Topup(ctx context.Context, req *entity.TransactionTopupRequest) (*entity.TransactionTopupResponse, error)
	Transfer(ctx context.Context, req *entity.TransactionTransferRequest) (*entity.TransactionTransferResponse, error)
	Withdraw(ctx context.Context, req *entity.TransactionWithdrawRequest) (*entity.TransactionWithdrawResponse, error)
}

type transactionUC struct {
	txWrapper           repository.TransactionWrapper
	userAccountRepo     repository.UserAccountRepo
	transactionRepo     repository.TransactionRepo
	balanceMovementRepo repository.BalanceMovementRepo
	providerSettingRepo repository.ProviderSettingRepo
	pinRepo             repository.PINRepo
}

func NewTransactionUC(
	txWrapper repository.TransactionWrapper,
	userAccountRepo repository.UserAccountRepo,
	transactionRepo repository.TransactionRepo,
	balanceMovementRepo repository.BalanceMovementRepo,
	providerSettingRepo repository.ProviderSettingRepo,
	pinRepo repository.PINRepo,
) TransactionUC {
	return &transactionUC{
		txWrapper:           txWrapper,
		userAccountRepo:     userAccountRepo,
		transactionRepo:     transactionRepo,
		balanceMovementRepo: balanceMovementRepo,
		providerSettingRepo: providerSettingRepo,
		pinRepo:             pinRepo,
	}
}

func (u *transactionUC) Topup(ctx context.Context, req *entity.TransactionTopupRequest) (*entity.TransactionTopupResponse, error) {
	account, err := u.userAccountRepo.GetUserAccountByVA(ctx, req.VANumber)
	if err != nil {
		return nil, err
	}

	providerSetting, err := u.providerSettingRepo.GetProviderSetting(ctx, constant.ProviderIDBankBCA, constant.TransactionTypeTopUp)
	if err != nil {
		return nil, err
	}

	if err = u.txWrapper.ExecuteTransaction(ctx,
		func(ctxTX context.Context) error {
			//update user account balance
			err = u.userAccountRepo.UpdateUserBalance(ctxTX, account.ID, req.Amount, constant.INCREASE)
			if err != nil {
				return err
			}

			//create new transaction
			newTransaction := &model.Transaction{}
			newTransaction.CreateNewTransaction(&entity.CreateNewTransaction{
				TransactionFrom:   constant.ProviderIDBankBCA,
				TransactionTo:     account.ID,
				ProviderID:        constant.ProviderIDBankBCA,
				ProviderSettingID: providerSetting.ID,
				Category:          constant.TransactionTypeTopUp,
				Amount:            req.Amount,
				AdminFee:          providerSetting.AdminFee,
				ProviderFee:       providerSetting.ProviderFee,
			})
			err = u.transactionRepo.CreateTransaction(ctxTX, newTransaction)
			if err != nil {
				return err
			}

			//create new balance movement
			newBalanceMovement := &model.BalanceMovement{}
			newBalanceMovement.CreateNewBalanceMovement(&entity.CreateNewBalanceMovement{
				UserAccountID: account.ID,
				TransactionID: newTransaction.ID,
				Cashflow:      constant.CashflowIN,
				Amount:        req.Amount,
				BalanceBefore: account.Balance,
				BalanceAfter:  account.Balance + req.Amount,
			})
			err = u.balanceMovementRepo.CreateBalanceMovement(ctxTX, newBalanceMovement)
			if err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &entity.TransactionTopupResponse{}, nil
}

func (u *transactionUC) Transfer(ctx context.Context, req *entity.TransactionTransferRequest) (*entity.TransactionTransferResponse, error) {
	verifiedPIN, err := u.pinRepo.GetVerifiedPINByTypeCache(ctx, req.AccountID, constant.PINTypeTransfer)
	if err != nil {
		return nil, err
	}
	if !verifiedPIN {
		return nil, utils.ErrForbidden("Forbidden to access this resource", "transactionUC.Transfer.verifiedPIN")
	}
	err = u.pinRepo.DeleteVerifiedPINByTypeCache(ctx, req.AccountID, constant.PINTypeTransfer)
	if err != nil {
		return nil, err
	}

	account, err := u.userAccountRepo.GetUserAccountByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	toAccount, err := u.userAccountRepo.GetUserAccountByID(ctx, req.ToAccountID)
	if err != nil {
		return nil, err
	}

	providerSetting, err := u.providerSettingRepo.GetProviderSetting(ctx, constant.ProviderIDMiraePay, constant.TransactionTypeTransfer)
	if err != nil {
		return nil, err
	}

	if req.Amount > account.Balance {
		return nil, utils.ErrBadRequest("Insufficient balance", "transactionUC.Transfer.InsufficientBalance")
	}

	if err = u.txWrapper.ExecuteTransaction(ctx,
		func(ctxTX context.Context) error {
			//update user account balance
			err = u.userAccountRepo.UpdateUserBalance(ctxTX, account.ID, req.Amount, constant.DECREASE)
			if err != nil {
				return err
			}

			err = u.userAccountRepo.UpdateUserBalance(ctxTX, toAccount.ID, req.Amount, constant.INCREASE)
			if err != nil {
				return err
			}

			//create new transaction
			newTransaction := &model.Transaction{}
			newTransaction.CreateNewTransaction(&entity.CreateNewTransaction{
				TransactionFrom:   account.ID,
				TransactionTo:     toAccount.ID,
				ProviderID:        constant.ProviderIDMiraePay,
				ProviderSettingID: providerSetting.ID,
				Category:          constant.TransactionTypeTransfer,
				Amount:            req.Amount,
				AdminFee:          providerSetting.AdminFee,
				ProviderFee:       providerSetting.ProviderFee,
			})
			err = u.transactionRepo.CreateTransaction(ctxTX, newTransaction)
			if err != nil {
				return err
			}

			//create new balance movement
			accountBalanceMovement := &model.BalanceMovement{}
			accountBalanceMovement.CreateNewBalanceMovement(&entity.CreateNewBalanceMovement{
				UserAccountID: account.ID,
				TransactionID: newTransaction.ID,
				Cashflow:      constant.CashflowOUT,
				Amount:        req.Amount,
				BalanceBefore: account.Balance,
				BalanceAfter:  account.Balance - req.Amount,
			})

			destinationBalanceMovement := &model.BalanceMovement{}
			destinationBalanceMovement.CreateNewBalanceMovement(&entity.CreateNewBalanceMovement{
				UserAccountID: toAccount.ID,
				TransactionID: newTransaction.ID,
				Cashflow:      constant.CashflowIN,
				Amount:        req.Amount,
				BalanceBefore: toAccount.Balance,
				BalanceAfter:  toAccount.Balance + req.Amount,
			})

			err = u.balanceMovementRepo.CreateBalanceMovement(ctxTX, accountBalanceMovement)
			if err != nil {
				return err
			}

			err = u.balanceMovementRepo.CreateBalanceMovement(ctxTX, destinationBalanceMovement)
			if err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &entity.TransactionTransferResponse{}, nil
}

func (u *transactionUC) Withdraw(ctx context.Context, req *entity.TransactionWithdrawRequest) (*entity.TransactionWithdrawResponse, error) {
	verifiedPIN, err := u.pinRepo.GetVerifiedPINByTypeCache(ctx, req.AccountID, constant.PINTypeWithdraw)
	if err != nil {
		return nil, err
	}
	if !verifiedPIN {
		return nil, utils.ErrForbidden("Forbidden to access this resource", "transactionUC.Withdraw.verifiedPIN")
	}
	err = u.pinRepo.DeleteVerifiedPINByTypeCache(ctx, req.AccountID, constant.PINTypeWithdraw)
	if err != nil {
		return nil, err
	}

	account, err := u.userAccountRepo.GetUserAccountByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	providerSetting, err := u.providerSettingRepo.GetProviderSetting(ctx, constant.ProviderIDBankBCA, constant.TransactionTypeWithdraw)
	if err != nil {
		return nil, err
	}

	if req.Amount > account.Balance {
		return nil, utils.ErrBadRequest("Insufficient balance", "transactionUC.Withdraw.InsufficientBalance")
	}

	if err = u.txWrapper.ExecuteTransaction(ctx,
		func(ctxTX context.Context) error {
			//update user account balance
			err = u.userAccountRepo.UpdateUserBalance(ctxTX, account.ID, req.Amount, constant.DECREASE)
			if err != nil {
				return err
			}

			//create new transaction
			newTransaction := &model.Transaction{}
			newTransaction.CreateNewTransaction(&entity.CreateNewTransaction{
				TransactionFrom:   constant.ProviderIDBankBCA,
				TransactionTo:     account.ID,
				ProviderID:        constant.ProviderIDBankBCA,
				ProviderSettingID: providerSetting.ID,
				Category:          constant.TransactionTypeWithdraw,
				Amount:            req.Amount,
				AdminFee:          providerSetting.AdminFee,
				ProviderFee:       providerSetting.ProviderFee,
			})
			err = u.transactionRepo.CreateTransaction(ctxTX, newTransaction)
			if err != nil {
				return err
			}

			//create new balance movement
			newBalanceMovement := &model.BalanceMovement{}
			newBalanceMovement.CreateNewBalanceMovement(&entity.CreateNewBalanceMovement{
				UserAccountID: account.ID,
				TransactionID: newTransaction.ID,
				Cashflow:      constant.CashflowOUT,
				Amount:        req.Amount,
				BalanceBefore: account.Balance,
				BalanceAfter:  account.Balance - req.Amount,
			})
			err = u.balanceMovementRepo.CreateBalanceMovement(ctxTX, newBalanceMovement)
			if err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &entity.TransactionWithdrawResponse{}, nil
}
