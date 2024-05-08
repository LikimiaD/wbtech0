package database

import (
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var CacheFilePath = filepath.Join("cache", "cache_backup.json")

func setupSignalHandler(orderService *OrderService) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		if err := orderService.SaveCacheToFile(CacheFilePath); err != nil {
			slog.Warn("Failed to save cache", "error", err.Error())
		}
		os.Exit(0)
	}()
}

func (s *OrderService) SaveCacheToFile(filename string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.Cache); err != nil {
		return err
	}
	return nil
}
