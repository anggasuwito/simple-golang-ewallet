package model

type UserAccount struct {
	BaseModel

	UserID  string `gorm:"column:user_id"`
	Balance int64  `gorm:"column:balance"`
	Status  string `gorm:"column:status"`
}
