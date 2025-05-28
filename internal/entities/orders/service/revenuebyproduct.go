package service

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/serror"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *service) RevenueByProduct(filter urlquery.DateRange) (
	models.RevenueByProductResp, *serror.ServiceError,
) {
	revenueByProduct, err := s.db.RevenueByProduct(filter)
	if err != nil {
		s.logger.Error("Error fetching revenue by product", zap.Error(err))
		return models.RevenueByProductResp{}, &serror.ServiceError{
			Code:  http.StatusInternalServerError,
			Error: errors.Wrap(err, "failed to fetch revenue by product"),
		}
	}
	return models.RevenueByProductResp{Products: revenueByProduct}, nil
}
