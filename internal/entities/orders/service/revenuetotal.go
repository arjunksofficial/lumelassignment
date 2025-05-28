package service

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/serror"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *service) RevenueTotal(
	filter urlquery.DateRange) (
	models.RevenueTotalResp, *serror.ServiceError,
) {
	totalRevenue, err := s.db.RevenueTotal(filter)
	if err != nil {
		s.logger.Error("Error fetching total revenue", zap.Error(err))
		return models.RevenueTotalResp{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to fetch total revenue"),
		}
	}
	return models.RevenueTotalResp{TotalRevenue: totalRevenue}, nil
}
