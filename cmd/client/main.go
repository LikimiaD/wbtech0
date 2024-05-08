package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/likimiad/wbtech0/internal/config"
	"github.com/likimiad/wbtech0/internal/data_generator"
	"github.com/likimiad/wbtech0/internal/handlers"
	"github.com/nats-io/stan.go"
	"log/slog"
)

func main() {
	cfg := config.GetConfig()

	sc, err := connectToNATS(cfg)
	if err != nil {
		handlers.FatalError("Failed to connect to NATS Streaming server", err)
	}
	defer sc.Close()

	ticker := time.NewTicker(time.Duration(cfg.NatsConfig.Ticker) * time.Second)
	defer ticker.Stop()

	slog.Info(fmt.Sprintf("Client will send requests every %d seconds", cfg.NatsConfig.Ticker))

	go runPublishingLoop(sc, ticker)

	waitForTerminateSignal()
	slog.Info("Shutting down gracefully...")
}

func connectToNATS(cfg *config.Config) (stan.Conn, error) {
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano())
	slog.Info("generate new client id", "client_id", clientID)

	sc, err := stan.Connect(cfg.NatsConfig.ClusterID, clientID, stan.NatsURL(cfg.NatsConfig.URL))
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func runPublishingLoop(sc stan.Conn, ticker *time.Ticker) {
	for range ticker.C {
		order := data_generator.GenerateRandomOrder()
		orderJSON, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			handlers.WarnError("error marshalling JSON", err)
			return
		}
		if err = sc.Publish("orders", orderJSON); err != nil {
			handlers.WarnError("error publishing message to 'orders'", err)
		} else {
			slog.Info("published message to 'orders'", "order", order)
		}
	}
	slog.Info("Ticker stopped, exiting publishing loop")
}

func waitForTerminateSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	slog.Info("Received terminate signal", "signal", sig)
}
