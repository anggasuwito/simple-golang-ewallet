package model

import "encoding/json"

type Transaction struct {
	BaseModel

	UserAccountID     string          `gorm:"column:user_account_id"`
	ProviderID        string          `gorm:"column:provider_id"`
	ProviderSettingID string          `gorm:"column:provider_setting_id"`
	ProviderName      string          `gorm:"column:provider_name"`
	ReferenceID       string          `gorm:"column:reference_id"`
	Category          string          `gorm:"column:category"`
	Amount            int64           `gorm:"column:amount"`
	AdminFee          int64           `gorm:"column:admin_fee"`
	ProviderFee       int64           `gorm:"column:provider_fee"`
	TotalAmount       int64           `gorm:"column:total_amount"`
	Status            string          `gorm:"column:status"`
	AdditionalInfo    json.RawMessage `gorm:"column:additional_info"`
}
