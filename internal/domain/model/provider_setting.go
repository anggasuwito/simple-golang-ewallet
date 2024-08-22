package model

func (m *ProviderSetting) TableName() string {
	return "provider_setting"
}

type ProviderSetting struct {
	BaseModel
	Provider Provider `gorm:"foreignKey:ProviderID"`

	ProviderID  string `gorm:"column:provider_id;size:36;index;"`
	Category    string `gorm:"column:category;size:255;"`
	MinAmount   int64  `gorm:"column:min_amount"`
	MaxAmount   int64  `gorm:"column:max_amount"`
	ProviderFee int64  `gorm:"column:provider_fee"`
	AdminFee    int64  `gorm:"column:admin_fee"`
}
