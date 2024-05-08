package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var (
	ErrBadRequest = errors.New("request failed validation")
)

func FatalError(text string, err error) {
	slog.Error(text, "error", err.Error())
	os.Exit(1)
}

func WarnError(text string, err error) {
	slog.Warn(text, "error", err.Error())
}

func FuncInfo(funcName string, start time.Time) {
	slog.Info(fmt.Sprintf("func '%s' completed the work in %v s", funcName, time.Since(start).Seconds()))
}
