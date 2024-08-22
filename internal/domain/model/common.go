package model

import "database/sql"

type BaseModel struct {
	ID        string       `gorm:"column:id;primaryKey;size:36;"`
	CreatedAt sql.NullTime `gorm:"column:created_at;autoCreateTime:milli"`
	CreatedBy string       `gorm:"column:created_by"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;autoUpdateTime:milli"`
	UpdatedBy string       `gorm:"column:updated_by"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
	DeletedBy string       `gorm:"column:deleted_by"`
	Note      string       `gorm:"column:note"`
}
