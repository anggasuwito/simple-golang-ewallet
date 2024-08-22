package model

func (m *UserAccount) TableName() string {
	return "user_account"
}

type UserAccount struct {
	BaseModel
	User User `gorm:"foreignKey:UserID"`

	UserID   string `gorm:"column:user_id;size:36;"`
	Balance  int64  `gorm:"column:balance"`
	Status   string `gorm:"column:status;size:100;"`
	VANumber string `gorm:"column:va_number;size:100;"`
	PIN      string `gorm:"column:pin;size:255;"`
}
