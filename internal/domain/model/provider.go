package model

type Provider struct {
	BaseModel

	Name   string `gorm:"column:name"`
	Status string `gorm:"column:status"`
}
