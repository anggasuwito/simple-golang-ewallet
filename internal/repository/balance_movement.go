package repository

import (
	"context"
	"gorm.io/gorm"
	"simple-golang-ewallet/internal/domain/model"
	"simple-golang-ewallet/internal/utils"
)

type BalanceMovementRepo interface {
	CreateBalanceMovement(ctx context.Context, data *model.BalanceMovement) error
}

type balanceMovementRepo struct {
	masterDB *gorm.DB
}

func NewBalanceMovementRepo(masterDB *gorm.DB) BalanceMovementRepo {
	return &balanceMovementRepo{
		masterDB: masterDB,
	}
}

func (r *balanceMovementRepo) useTX(ctx context.Context) *gorm.DB {
	if tx := utils.GetTransactionFromContext(ctx); tx != nil {
		return tx
	}
	return r.masterDB
}

func (r *balanceMovementRepo) CreateBalanceMovement(ctx context.Context, data *model.BalanceMovement) error {
	db := r.useTX(ctx)
	err := db.Debug().Create(data).Error
	if err != nil {
		return utils.ErrInternal("Failed create new balance movement : "+err.Error(), "balanceMovementRepo.CreateBalanceMovement")
	}
	return nil
}
