package handlers

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/pkg/core/responsehelper"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) RevenueByProduct(c *gin.Context) {
	dateRange, sErr := urlquery.GetDateRangeQuery(c)
	if sErr != nil {
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	resp, sErr := h.svc.RevenueByProduct(dateRange)
	if sErr != nil {
		if sErr.Code >= http.StatusInternalServerError {
			h.logger.Error("Error fetching revenue by product", zap.Error(sErr.Error))
			c.JSON(sErr.Code, responsehelper.NewCommonResponse("Internal Server Error"))
			return
		}
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
