package model

type User struct {
	BaseModel

	Phone    string `gorm:"column:phone"`
	VANumber string `gorm:"column:va_number"`
	PIN      string `gorm:"column:pin"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
}
