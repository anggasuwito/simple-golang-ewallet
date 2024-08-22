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

func GenerateJWT(user *model.User) (entity.JWT, error) {
	tokenB := jwt.New(jwt.SigningMethodHS256)
	claims := tokenB.Claims.(jwt.MapClaims)
	cfg := config.GetConfig()
	tokenExpiredDuration, _ := time.ParseDuration(cfg.AccessTokenExpireDuration)

	// Set payload
	id := uuid.New().String()
	expiredAt := TimeNow().Add(tokenExpiredDuration).Unix()
	claims["expired_at"] = expiredAt
	claims[claimsDataKey] = entity.JWTClaim{
		ID:        "",
		ExpiredAt: expiredAt,
		UserInfo:  entity.JWTClaimUserInfo{},
	}

			entity.JWTClaimUserInfo{
		ID:    user.ID,
		Phone: user.Phone,
		Name:  user.Name,
		Email: user.Email,
	}

	tokenString, err := tokenB.SignedString([]byte(cfg.AccessTokenSecret))
	if err != nil {
		return entity.JWT{}, err
	}
	return entity.JWT{
		ID:    id,
		Token: tokenString,
	}, nil
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

	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		return jwtClaim, err
	}

	err = json.Unmarshal(jsonClaims, &jwtClaim)
	if err != nil {
		return jwtClaim, err
	}
	return jwtClaim, nil
}
