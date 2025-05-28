package handlers

import (
	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/service"
	"github.com/arjunksofficial/lumelassignment/pkg/core/logger"
	"go.uber.org/zap"
)

type Handler struct {
	svc    service.Service
	logger *zap.Logger
}

func New() *Handler {
	return &Handler{
		svc:    service.GetService(),
		logger: logger.GetLogger(),
	}
}
