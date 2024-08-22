package repository

import (
	"context"
	"gorm.io/gorm"
	"simple-golang-ewallet/internal/domain/model"
	"simple-golang-ewallet/internal/utils"
)

type UserAccountRepo interface {
	GetUserAccount(ctx context.Context, phone string) (*model.UserAccount, error)
}

type userAccountRepo struct {
	masterDB *gorm.DB
}

func NewUserAccountRepo(masterDB *gorm.DB) UserAccountRepo {
	return &userAccountRepo{
		masterDB: masterDB,
	}
}

func (r *userAccountRepo) GetUserAccount(ctx context.Context, phone string) (*model.UserAccount, error) {
	var data model.UserAccount

	err := r.masterDB.
		Debug().
		Model(&model.UserAccount{}).
		Where("deleted_at IS NULL").
		Where("phone = ?", phone).
		First(&data).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrNotFound("User Account Not Found", "userAccountRepo.GetUserAccount.ErrRecordNotFound")
		}
		return nil, utils.ErrInternal("Failed get user account : "+err.Error(), "userAccountRepo.GetUserAccount")
	}

	return &data, nil
}
