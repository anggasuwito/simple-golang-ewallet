package entity

type (
	TransactionTopupRequest struct {
		VANumber string `json:"va_number"`
		Amount   int64  `json:"amount"`
	}

	TransactionTopupResponse struct {
	}

	TransactionTransferRequest struct {
	}

	TransactionTransferResponse struct {
	}

	TransactionWithdrawRequest struct {
	}

	TransactionWithdrawResponse struct {
	}

	CreateNewTransaction struct {
		TransactionFrom   string
		TransactionTo     string
		ProviderID        string
		ProviderSettingID string
		ProviderName      string
		Category          string
		Source            string
		Amount            int64
		AdminFee          int64
		ProviderFee       int64
		AdditionalInfo    []byte
	}
)
