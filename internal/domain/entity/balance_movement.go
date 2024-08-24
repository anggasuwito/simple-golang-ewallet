package entity

type (
	BalanceMovement struct {
		ID            string `json:"id"`
		UserAccountID string `json:"user_account_id"`
		TransactionID string `json:"transaction_id"`
		Cashflow      string `json:"cashflow"`
		Amount        int64  `json:"amount"`
		Description   string `json:"description"`
	}

	CreateNewBalanceMovement struct {
		UserAccountID string
		TransactionID string
		Cashflow      string
		Amount        int64
		Description   string
		BalanceBefore int64
		BalanceAfter  int64
	}
)
