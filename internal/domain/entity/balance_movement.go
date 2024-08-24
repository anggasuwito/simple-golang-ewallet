package entity

type (
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
