package repository

import (
	"context"
	"gorm.io/gorm"
	"simple-golang-ewallet/internal/domain/model"
	"simple-golang-ewallet/internal/utils"
)

type UserAccountRepo interface {
	GetUserAccountByPhone(ctx context.Context, phone string) (*model.UserAccount, error)
	GetUserAccountByID(ctx context.Context, id string) (*model.UserAccount, error)
}

type userAccountRepo struct {
	masterDB *gorm.DB
}

func NewUserAccountRepo(masterDB *gorm.DB) UserAccountRepo {
	return &userAccountRepo{
		masterDB: masterDB,
	}
}

func (r *userAccountRepo) GetUserAccountByPhone(ctx context.Context, phone string) (*model.UserAccount, error) {
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
			return nil, utils.ErrNotFound("User Account Not Found", "userAccountRepo.GetUserAccountByPhone.ErrRecordNotFound")
		}
		return nil, utils.ErrInternal("Failed get user account : "+err.Error(), "userAccountRepo.GetUserAccountByPhone")
	}

	return &data, nil
}

func (r *userAccountRepo) GetUserAccountByID(ctx context.Context, id string) (*model.UserAccount, error) {
	var data model.UserAccount

	err := r.masterDB.
		Debug().
		Model(&model.UserAccount{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&data).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrNotFound("User Account Not Found", "userAccountRepo.GetUserAccountByID.ErrRecordNotFound")
		}
		return nil, utils.ErrInternal("Failed get user account : "+err.Error(), "userAccountRepo.GetUserAccountByID")
	}

	return &data, nil
}
