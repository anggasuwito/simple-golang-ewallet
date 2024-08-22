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

type UserHandler struct {
	userUC usecase.UserUC
}

func NewUserHandler(
	userUC usecase.UserUC,
) *UserHandler {
	return &UserHandler{
		userUC: userUC,
	}
}

func (h *UserHandler) SetupHandlers(r *gin.Engine) {
	userPathV1 := r.Group("/v1")
	userPathV1.Use(middleware.TokenChecker)
	userPathV1.GET("/user/get-info", h.getInfo)
	userPathV1.GET("/user/get-transaction-history", h.getTransactionHistory)
}

func (h *UserHandler) getInfo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.userUC.GetInfo(ctx, &entity.UserGetInfoRequest{})
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *UserHandler) getTransactionHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.userUC.GetTransactionHistory(ctx, &entity.UserGetTransactionHistoryRequest{})
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
