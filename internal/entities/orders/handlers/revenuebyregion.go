package handlers

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/responsehelper"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	RevenueByRegion give stats of revenue by each region
//
// @Summary		stats of revenue by each region
// @Description	get revenue by region for a given date range
// @ID			revenue-by-region
// @Tags		revenue
// @Produce		json
// @Param		start_date	 query		string                   	  false	"Start Date"
// @Param		end_date	 query		string                   	  false	"End Date"
// @Success		200	         {object}	models.RevenueByRegionResp
// @Failure		400	         {object}	responsehelper.CommonResponse
// @Failure		500	         {object}	responsehelper.CommonResponse
// @Router		/api/v1/orders/revenue-by-region [get]
func (h *Handler) RevenueByRegion(c *gin.Context) {
	dateRange, sErr := urlquery.GetDateRangeQuery(c)
	if sErr != nil {
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	var resp models.RevenueByRegionResp
	resp, sErr = h.svc.RevenueByRegion(dateRange)
	if sErr != nil {
		if sErr.Code >= http.StatusInternalServerError {
			h.logger.Error("Error fetching revenue by region", zap.Error(sErr.Error))
			c.JSON(sErr.Code, responsehelper.NewCommonResponse("Internal Server Error"))
			return
		}
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
