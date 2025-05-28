package service

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/serror"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// RevenueByRegion retrieves revenue data grouped by region within a specified date range.
func (s *service) RevenueByRegion(
	filter urlquery.DateRange) (
	models.RevenueByRegionResp, *serror.ServiceError,
) {
	revenueByRegion, err := s.db.RevenueByRegion(filter)
	if err != nil {
		s.logger.Error("Error fetching revenue by region", zap.Error(err))
		return models.RevenueByRegionResp{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to fetch revenue by region"),
		}
	}
	return models.RevenueByRegionResp{Regions: revenueByRegion}, nil
}
