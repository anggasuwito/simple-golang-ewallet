package model

func (m *Provider) TableName() string {
	return "provider"
}

type Provider struct {
	BaseModel
	ProviderSetting []ProviderSetting `gorm:"foreignKey:ProviderID"`

	Name   string `gorm:"column:name"`
	Status string `gorm:"column:status"`
}
