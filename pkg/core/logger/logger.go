package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger == nil {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			panic(fmt.Sprintf("failed to create logger: %v", err))
		}
	}
	return logger
}
func CloseLogger() {
	if logger != nil {
		err := logger.Sync()
		if err != nil {
			fmt.Printf("failed to close logger: %v\n", err)
		}
	}
}
