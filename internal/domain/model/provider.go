package model

func (m *Provider) TableName() string {
	return "provider"
}

type Provider struct {
	BaseModel
	ProviderSetting []ProviderSetting `gorm:"foreignKey:ProviderID"`

	Name           string `gorm:"column:name;size:255;"`
	IdentityNumber string `gorm:"column:identity_number;size:100"`
	Status         string `gorm:"column:status;size:255;"`
}
