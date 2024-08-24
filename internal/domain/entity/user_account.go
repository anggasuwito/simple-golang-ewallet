package entity

type (
	UserAccount struct {
		ID       string `json:"id"`
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Balance  int64  `json:"balance"`
		Status   string `json:"status"`
		VANumber string `json:"va_number"`
	}

	UserAccGetInfoRequest struct {
		AccountID string
	}

	UserAccGetInfoResponse struct {
		*UserAccount
	}

	UserAccGetTransactionHistoryRequest struct {
		AccountID string
	}

	UserAccGetTransactionHistoryResponse struct {
		List []*BalanceMovement `json:"list"`
	}
)
