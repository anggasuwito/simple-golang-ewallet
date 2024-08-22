package model

func (m *Provider) TableName() string {
	return "provider"
}

type Provider struct {
	BaseModel
	ProviderSetting []ProviderSetting `gorm:"foreignKey:ProviderID"`

	Name   string `gorm:"column:name;size:255;"`
	Status string `gorm:"column:status;size:255;"`
}
