package repository

import (
	"context"
	"gorm.io/gorm"
	"simple-golang-ewallet/internal/domain/model"
	"simple-golang-ewallet/internal/utils"
)

type ProviderSettingRepo interface {
	GetSettingByProviderID(ctx context.Context, providerID string) (*model.ProviderSetting, error)
}

type providerSettingRepo struct {
	masterDB *gorm.DB
}

func NewProviderSettingRepo(masterDB *gorm.DB) ProviderSettingRepo {
	return &providerSettingRepo{
		masterDB: masterDB,
	}
}

func (r *providerSettingRepo) GetSettingByProviderID(ctx context.Context, providerID string) (*model.ProviderSetting, error) {
	var data model.ProviderSetting

	err := r.masterDB.
		Debug().
		Model(&model.ProviderSetting{}).
		Where("deleted_at IS NULL").
		Where("provider_id = ?", providerID).
		First(&data).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrNotFound("Provider Setting Not Found", "providerSettingRepo.GetSettingByProviderID.ErrRecordNotFound")
		}
		return nil, utils.ErrInternal("Failed get provider setting : "+err.Error(), "providerSettingRepo.GetSettingByProviderID")
	}

	return &data, nil
}
