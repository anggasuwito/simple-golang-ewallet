package model

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
