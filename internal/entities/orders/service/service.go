package service

import (
	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/store"
	"github.com/arjunksofficial/lumelassignment/pkg/core/logger"
	"github.com/arjunksofficial/lumelassignment/pkg/core/serror"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"go.uber.org/zap"
)

type service struct {
	db     store.Store
	logger *zap.Logger
}

type Service interface {
	TriggerImport() error
	RevenueTotal(filter urlquery.DateRange) (models.RevenueTotalResp, *serror.ServiceError)
	RevenueByProduct(filter urlquery.DateRange) (models.RevenueByProductResp, *serror.ServiceError)
	RevenueByCategory(filter urlquery.DateRange) (models.RevenueByCategoryResp, *serror.ServiceError)
	RevenueByRegion(filter urlquery.DateRange) (models.RevenueByRegionResp, *serror.ServiceError)
	RevenueByTrends(filter urlquery.DateRange) (models.RevenueByTrendsResp, *serror.ServiceError)
}

func GetService() Service {
	return &service{
		db:     store.GetStore(),
		logger: logger.GetLogger(),
	}
}
