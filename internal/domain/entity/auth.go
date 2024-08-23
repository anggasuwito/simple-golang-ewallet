package entity

type (
	JWTClaimAccountInfo struct {
		ID    string `json:"id"`
		Phone string `json:"phone"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	JWTClaim struct {
		ID          string              `json:"id"`
		ExpiredAt   int64               `json:"expired_at"`
		AccountInfo JWTClaimAccountInfo `json:"account_info"`
	}

	AuthLoginPINRequest struct {
		Phone string `json:"phone"`
		PIN   string `json:"pin"`
	}

	AuthLoginPINResponse struct {
		AccessToken string `json:"access_token"`
	}

	AuthVerifyPINRequest struct {
		AccountID string
		PIN       string `json:"pin"`
		Type      string `json:"type"`
	}

	AuthVerifyPINResponse struct {
	}
)
