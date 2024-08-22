package model

import "encoding/json"

func (m *Transaction) TableName() string {
	return "transaction"
}

type Transaction struct {
	BaseModel
	UserAccount     UserAccount     `gorm:"foreignKey:UserAccountID"`
	Provider        Provider        `gorm:"foreignKey:ProviderID"`
	ProviderSetting ProviderSetting `gorm:"foreignKey:ProviderSettingID"`

	UserAccountID     string          `gorm:"column:user_account_id;size:36;"`
	ProviderID        string          `gorm:"column:provider_id;size:36;"`
	ProviderSettingID string          `gorm:"column:provider_setting_id;size:36;"`
	ProviderName      string          `gorm:"column:provider_name;size:255;"`
	ReferenceID       string          `gorm:"column:reference_id;size:255;"`
	Category          string          `gorm:"column:category;size:255;"`
	Amount            int64           `gorm:"column:amount"`
	AdminFee          int64           `gorm:"column:admin_fee"`
	ProviderFee       int64           `gorm:"column:provider_fee"`
	TotalAmount       int64           `gorm:"column:total_amount"`
	Status            string          `gorm:"column:status;size:100;"`
	AdditionalInfo    json.RawMessage `gorm:"column:additional_info;type:jsonb"`
}
