package model

func (m *UserAccount) TableName() string {
	return "user_account"
}

type UserAccount struct {
	BaseModel
	User User `gorm:"foreignKey:UserID"`

	UserID  string `gorm:"column:user_id"`
	Balance int64  `gorm:"column:balance"`
	Status  string `gorm:"column:status"`
}
