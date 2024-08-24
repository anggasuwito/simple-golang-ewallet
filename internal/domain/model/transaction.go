package model

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"simple-golang-ewallet/internal/domain/entity"
	"time"
)

func (m *Transaction) TableName() string {
	return "transaction"
}

type Transaction struct {
	BaseModel
	Provider        Provider        `gorm:"foreignKey:ProviderID"`
	ProviderSetting ProviderSetting `gorm:"foreignKey:ProviderSettingID"`

	TransactionFrom   string          `gorm:"column:transaction_from;size:255;"`
	TransactionTo     string          `gorm:"column:transaction_to;size:255;"`
	ProviderID        string          `gorm:"column:provider_id;size:36;"`
	ProviderSettingID string          `gorm:"column:provider_setting_id;size:36;"`
	ProviderName      string          `gorm:"column:provider_name;size:255;"`
	ReferenceID       string          `gorm:"column:reference_id;size:255;"`
	Category          string          `gorm:"column:category;size:255;"`
	Source            string          `gorm:"column:source;size:255;"`
	Amount            int64           `gorm:"column:amount"`
	AdminFee          int64           `gorm:"column:admin_fee"`
	ProviderFee       int64           `gorm:"column:provider_fee"`
	TotalAmount       int64           `gorm:"column:total_amount"`
	Status            string          `gorm:"column:status;size:100;"`
	ReceivedAt        sql.NullTime    `gorm:"column:received_at"`
	FailedAt          sql.NullTime    `gorm:"column:failed_at"`
	Description       string          `gorm:"column:description;size:255;"`
	AdditionalInfo    json.RawMessage `gorm:"column:additional_info;type:jsonb"`
}

func (m *Transaction) CreateNewTransaction(req *entity.CreateNewTransaction) {
	m.ID = uuid.New().String()
	m.TransactionFrom = req.TransactionFrom
	m.TransactionTo = req.TransactionTo
	m.ProviderID = req.ProviderID
	m.ProviderSettingID = req.ProviderSettingID
	m.ProviderName = req.ProviderName
	m.ReferenceID = uuid.New().String()
	m.Category = req.Category
	m.Amount = req.Amount
	m.AdminFee = req.AdminFee
	m.ProviderFee = req.ProviderFee
	m.TotalAmount = req.Amount + req.AdminFee + req.ProviderFee
	m.Status = "SUCCESS"
	m.ReceivedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	m.AdditionalInfo = req.AdditionalInfo
}
