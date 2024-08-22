package model

type ProviderSetting struct {
	BaseModel

	ProviderID  string `gorm:"column:provider_id"`
	Category    string `gorm:"column:category"`
	MinAmount   int64  `gorm:"column:min_amount"`
	MaxAmount   int64  `gorm:"column:max_amount"`
	ProviderFee int64  `gorm:"column:provider_fee"`
	AdminFee    int64  `gorm:"column:admin_fee"`
}
