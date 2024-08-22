package model

type Transaction struct {
	BaseModel

	ReferenceID     string `gorm:"column:reference_id"`
	ReferenceSource string `gorm:"column:reference_source"`
	TransactionFrom string `gorm:"column:transaction_from"`
	TransactionTo   string `gorm:"column:transaction_to"`
	Amount          int64  `gorm:"column:amount"`
	Type            string `gorm:"column:type"`
	CashFlow        string `gorm:"column:cash_flow"`
	Status          string `gorm:"column:status"`
	AdditionalInfo  string `gorm:"column:additional_info"`
}
