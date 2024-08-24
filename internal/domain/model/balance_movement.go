package model

import (
	"github.com/google/uuid"
	"simple-golang-ewallet/internal/domain/entity"
)

func (m *BalanceMovement) TableName() string {
	return "balance_movement"
}

type BalanceMovement struct {
	BaseModel

	UserAccountID string `gorm:"column:user_account_id;size:36;"`
	TransactionID string `gorm:"column:transaction_id;size:36;"`
	Cashflow      string `gorm:"column:cashflow;size:1m.0;"`
	Amount        int64  `gorm:"column:amount"`
	Description   string `gorm:"column:description;size:255;"`
	BalanceBefore int64  `gorm:"column:balance_before"`
	BalanceAfter  int64  `gorm:"column:balance_after"`
}

func (m *BalanceMovement) CreateNewBalanceMovement(req *entity.CreateNewBalanceMovement) {
	m.ID = uuid.New().String()
	m.UserAccountID = req.UserAccountID
	m.TransactionID = req.TransactionID
	m.Cashflow = req.Cashflow
	m.Amount = req.Amount
	m.Description = req.Description
	m.BalanceBefore = req.BalanceBefore
	m.BalanceAfter = req.BalanceAfter
}

func (m *BalanceMovement) ToEntity() *entity.BalanceMovement {
	return &entity.BalanceMovement{
		ID:            m.ID,
		UserAccountID: m.UserAccountID,
		TransactionID: m.TransactionID,
		Cashflow:      m.Cashflow,
		Amount:        m.Amount,
		Description:   m.Description,
	}
}
