package model

func (m *User) TableName() string {
	return "user"
}

type User struct {
	BaseModel
	UserAccount []UserAccount `gorm:"foreignKey:UserID"`

	Phone string `gorm:"column:phone;size:100;"`
	Name  string `gorm:"column:name;size:255;"`
	Email string `gorm:"column:email;size:255;"`
}
