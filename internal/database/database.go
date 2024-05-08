package database

import (
	"errors"
	"fmt"
	"github.com/likimiad/wbtech0/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"log/slog"
	"time"
)

type Database struct {
	*gorm.DB
}

var (
	ErrConnect   = errors.New("failed to connect to the database")
	ErrMigration = errors.New("error during database auto-migration")
)

func makeConnection(cfg config.DatabaseConfig) *Database {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		cfg.Host, cfg.User, cfg.Name, cfg.Password, cfg.Port)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error("error when connecting to the database", "error", err)
		return nil
	}
	return &Database{gormDB}
}

func GetDB(cfg config.DatabaseConfig) (*OrderService, error) {
	defer func(start time.Time) {
		slog.Info("database connection is established", "duration", time.Since(start))
	}(time.Now())

	db := makeConnection(cfg)
	if db == nil {
		return nil, ErrConnect
	}

	if err := db.AutoMigrate(&Order{}, &Delivery{}, &Payment{}, &Item{}); err != nil {
		slog.Error("error during database auto-migration", "error", err)
		return nil, ErrMigration
	}

	orderService := NewOrderService(db)
	if err := orderService.RestoreCache(); err != nil {
		log.Fatalf("Failed to restore cache: %v", err)
	}
	go setupSignalHandler(orderService)

	return orderService, nil
}
