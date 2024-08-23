package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"simple-golang-ewallet/config"
	handler "simple-golang-ewallet/internal/handler/http"
	"simple-golang-ewallet/internal/repository"
	"simple-golang-ewallet/internal/usecase"
	"simple-golang-ewallet/internal/utils"
	"syscall"
	"time"
)

func setupMiddlewares(r *gin.Engine) {
	//r.Use(cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func setupRouters(r *gin.Engine) {
	cfg := config.GetConfig()

	r.GET("", func(c *gin.Context) {
		utils.ResponseSuccess(c, "Success run "+cfg.AppVersion, nil)
	})

	userAccountRepo := repository.NewUserAccountRepo(cfg.DBMaster)
	transactionRepo := repository.NewTransactionRepo(cfg.DBMaster)
	pinRepo := repository.NewPINRepo(cfg.RedisClient)

	authUC := usecase.NewAuthUC(userAccountRepo, pinRepo)
	userAccountUC := usecase.NewUserAccUC(userAccountRepo)
	transactionUC := usecase.NewTransactionUC(transactionRepo)

	handler.NewAuthHandler(authUC).SetupHandlers(r)
	handler.NewUserAccHandler(userAccountUC).SetupHandlers(r)
	handler.NewTransactionHandler(transactionUC).SetupHandlers(r)
}

func StartHttpServer() {
	r := gin.New()
	cfg := config.GetConfig()
	setupMiddlewares(r)
	setupRouters(r)
	addr := fmt.Sprintf("%s:%s", cfg.HttpHost, cfg.HttpPort)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		log.Println("[api] http server has been started in " + addr)
		log.Fatal(srv.ListenAndServe().Error())
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("[api] server forced to shutdown " + err.Error())
	}
}
