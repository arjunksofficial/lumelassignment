package service

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/serror"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *service) RevenueByTrends(
	filter urlquery.DateRange) (
	models.RevenueByTrendsResp, *serror.ServiceError,
) {
	revenueByTrends, err := s.db.RevenueByTrends(filter)
	if err != nil {
		s.logger.Error("Error fetching revenue by trends", zap.Error(err))
		return models.RevenueByTrendsResp{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to fetch revenue by trends"),
		}
	}
	return models.RevenueByTrendsResp{Trends: revenueByTrends}, nil
}
