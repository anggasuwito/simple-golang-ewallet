package repository

import (
	"gorm.io/gorm"
)

type UserRepo interface {
}

type userRepo struct {
	masterDB *gorm.DB
}

func NewUserRepo(masterDB *gorm.DB) UserRepo {
	return &userRepo{
		masterDB: masterDB,
	}
}
