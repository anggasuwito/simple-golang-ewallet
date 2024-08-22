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

type UserAccountHandler struct {
	userAccUC usecase.UserAccUC
}

func NewUserAccHandler(
	userAccUC usecase.UserAccUC,
) *UserAccountHandler {
	return &UserAccountHandler{
		userAccUC: userAccUC,
	}
}

func (h *UserAccountHandler) SetupHandlers(r *gin.Engine) {
	userAccPathV1 := r.Group("/v1")
	userAccPathV1.Use(middleware.TokenChecker)
	userAccPathV1.GET("/user-account/get-info", h.getInfo)
	userAccPathV1.GET("/user-account/get-transaction-history", h.getTransactionHistory)
}

func (h *UserAccountHandler) getInfo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.userAccUC.GetInfo(ctx, &entity.UserAccGetInfoRequest{})
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *UserAccountHandler) getTransactionHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.userAccUC.GetTransactionHistory(ctx, &entity.UserAccGetTransactionHistoryRequest{})
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
