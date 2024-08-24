package model

import "simple-golang-ewallet/internal/domain/entity"

func (m *UserAccount) TableName() string {
	return "user_account"
}

type UserAccount struct {
	BaseModel

	Phone    string `gorm:"column:phone;size:100;"`
	Name     string `gorm:"column:name;size:255;"`
	Email    string `gorm:"column:email;size:255;"`
	Balance  int64  `gorm:"column:balance"`
	Status   string `gorm:"column:status;size:100;"`
	VANumber string `gorm:"column:va_number;unique;size:100;"`
	PIN      string `gorm:"column:pin;size:255;"`
}

func (m *UserAccount) ToEntity() *entity.UserAccount {
	return &entity.UserAccount{
		ID:       m.ID,
		Phone:    m.Phone,
		Name:     m.Name,
		Email:    m.Email,
		Balance:  m.Balance,
		Status:   m.Status,
		VANumber: m.VANumber,
	}
}
