package model

func (m *User) TableName() string {
	return "user"
}

type User struct {
	BaseModel
	UserAccount []UserAccount `gorm:"foreignKey:UserID"`

	Phone    string `gorm:"column:phone;size:100;"`
	VANumber string `gorm:"column:va_number;size:100;"`
	PIN      string `gorm:"column:pin;size:255;"`
	Name     string `gorm:"column:name;size:255;"`
	Email    string `gorm:"column:email;size:255;"`
}
