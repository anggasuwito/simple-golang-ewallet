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

type TransactionHandler struct {
	transactionUC usecase.TransactionUC
}

func NewTransactionHandler(
	transactionUC usecase.TransactionUC,
) *TransactionHandler {
	return &TransactionHandler{
		transactionUC: transactionUC,
	}
}

func (h *TransactionHandler) SetupHandlers(r *gin.Engine) {
	transactionPathV1 := r.Group("/v1")
	transactionPathV1.Use(middleware.TokenChecker)
	transactionPathV1.POST("/transaction/topup", h.topup)
	transactionPathV1.POST("/transaction/transfer", h.transfer)
	transactionPathV1.POST("/transaction/withdraw", h.withdraw)
}

func (h *TransactionHandler) topup(c *gin.Context) {
	var req entity.TransactionTopupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "TransactionHandler.topup.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.transactionUC.Topup(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *TransactionHandler) transfer(c *gin.Context) {
	var req entity.TransactionTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "TransactionHandler.transfer.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.transactionUC.Transfer(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *TransactionHandler) withdraw(c *gin.Context) {
	var req entity.TransactionWithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "TransactionHandler.withdraw.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.transactionUC.Withdraw(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
