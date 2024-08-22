package model

func (m *User) TableName() string {
	return "user"
}

type User struct {
	BaseModel
	UserAccount []UserAccount `gorm:"foreignKey:UserID"`

	Phone    string `gorm:"column:phone"`
	VANumber string `gorm:"column:va_number"`
	PIN      string `gorm:"column:pin"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
}
