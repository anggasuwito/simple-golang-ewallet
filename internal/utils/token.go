package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"simple-golang-ewallet/config"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/domain/model"
	"time"
)

const claimsDataKey = "claims_data"

func GenerateJWT(acc *model.UserAccount) (tokenStr string, data entity.JWTClaim, err error) {
	tokenB := jwt.New(jwt.SigningMethodHS256)
	claims := tokenB.Claims.(jwt.MapClaims)
	cfg := config.GetConfig()
	tokenExpiredDuration, _ := time.ParseDuration(cfg.AccessTokenExpireDuration)

	// Set payload
	expiredAt := TimeNow().Add(tokenExpiredDuration).Unix()
	data = entity.JWTClaim{
		ID:        uuid.New().String(),
		ExpiredAt: expiredAt,
		AccountInfo: entity.JWTClaimAccountInfo{
			ID:    acc.ID,
			Phone: acc.Phone,
			Name:  acc.Name,
			Email: acc.Email,
		},
	}

	claims["expired_at"] = expiredAt
	claims[claimsDataKey] = data

	tokenStr, err = tokenB.SignedString([]byte(cfg.AccessTokenSecret))
	if err != nil {
		return "", data, err
	}
	return tokenStr, data, err
}

func VerifyJWT(token string) (jwtClaim entity.JWTClaim, err error) {
	cfg := config.GetConfig()
	secret := cfg.AccessTokenSecret

	tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return jwtClaim, err
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return jwtClaim, err
	}

	jsonClaims, err := json.Marshal(claims[claimsDataKey])
	if err != nil {
		return jwtClaim, err
	}

	err = json.Unmarshal(jsonClaims, &jwtClaim)
	if err != nil {
		return jwtClaim, err
	}
	return jwtClaim, nil
}
