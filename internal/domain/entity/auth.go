package entity

type (
	JWTClaimUserInfo struct {
		ID    string `json:"id"`
		Phone string `json:"phone"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	JWTClaim struct {
		ID        string           `json:"id"`
		ExpiredAt int64            `json:"expired_at"`
		UserInfo  JWTClaimUserInfo `json:"user_info"`
	}

	AuthLoginPINRequest struct {
	}

	AuthLoginPINResponse struct {
	}

	AuthVerifyPINRequest struct {
	}

	AuthVerifyPINResponse struct {
	}
)
