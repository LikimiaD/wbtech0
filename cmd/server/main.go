package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/likimiad/wbtech0/api"
	"github.com/likimiad/wbtech0/internal/config"
	"github.com/likimiad/wbtech0/internal/database"
	"github.com/likimiad/wbtech0/internal/handlers"
	"github.com/likimiad/wbtech0/internal/validator"
	"github.com/nats-io/stan.go"
	"log/slog"
)

// @title WBTech Task0 Go API
// @version 0.0.1
// @description API Server for demonstrate work

func main() {
	cfg := config.GetConfig()

	db, err := database.GetDB(cfg.DatabaseConfig)
	if err != nil {
		handlers.FatalError("failed to connect to database", err)
	}
	slog.Info("successfully connected to the database")

	server := api.NewServer(db)
	go func() {
		if err := server.Start(cfg.HTTPServer.Address); err != nil {
			handlers.FatalError("failed to start server", err)
		}
	}()

	sc, err := stan.Connect(cfg.NatsConfig.ClusterID, cfg.NatsConfig.ClientID, stan.NatsURL(cfg.NatsConfig.URL))
	if err != nil {
		handlers.FatalError("failed to connect to NATS Streaming", err)
	}
	defer func() {
		if err := sc.Close(); err != nil {
			handlers.WarnError("error while closing NATS Connection", err)
		}
	}()

	sub, err := sc.Subscribe("orders", func(msg *stan.Msg) {
		defer func(start time.Time) {
			handlers.FuncInfo("orders subscribe", start)
		}(time.Now())

		var order database.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			handlers.WarnError("error unmarshalling message", err)
			return
		}
		if !validator.IsValidOrder(&order) {
			handlers.WarnError("bad request", handlers.ErrBadRequest)
			return
		}
		if err := db.CreateOrder(&order); err != nil {
			handlers.WarnError("Failed to save order to database", err)
			return
		}

		slog.Info("order processed and saved", "uid", order.OrderUID)
	})
	if err != nil {
		handlers.FatalError("Failed to subscribe to channel 'orders'", err)
	}
	defer func() {
		if err := sub.Unsubscribe(); err != nil {
			handlers.WarnError("error while unsubscribe on NATS connection", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	slog.Info("Received terminate signal", "signal", sig)

	slog.Info("Shutting down gracefully...")
}
