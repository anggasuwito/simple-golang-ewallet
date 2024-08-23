package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"simple-golang-ewallet/internal/domain/entity"
	"simple-golang-ewallet/internal/handler/http/middleware"
	"simple-golang-ewallet/internal/usecase"
	"simple-golang-ewallet/internal/utils"
	"time"
)

type AuthHandler struct {
	authUC usecase.AuthUC
}

func NewAuthHandler(
	authUC usecase.AuthUC,
) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
	}
}

func (h *AuthHandler) SetupHandlers(r *gin.Engine) {
	authPathV1 := r.Group("/v1")
	authPathV1.POST("/auth/login-pin", h.loginPIN)
	authPathV1.POST("/auth/verify-pin", middleware.TokenChecker, h.verifyPIN)
}

func (h *AuthHandler) loginPIN(c *gin.Context) {
	var req entity.AuthLoginPINRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "AuthHandler.loginPIN.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.authUC.LoginPIN(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *AuthHandler) verifyPIN(c *gin.Context) {
	var req entity.AuthVerifyPINRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "AuthHandler.verifyPIN.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	ctxVal := middleware.GetContextValue(c)
	req.AccountID = ctxVal.AccountInfo.ID
	resp, err := h.authUC.VerifyPIN(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
