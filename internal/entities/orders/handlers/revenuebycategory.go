package handlers

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/responsehelper"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	RevenueByCategory give stats of revenue by each category
//
// @Summary		stats of revenue by each category
// @Description	get revenue by category for a given date range
// @ID			revenue-by-category
// @Tags		revenue
// @Produce		json
// @Param		start_date	 query		string                   	  false	"Start Date"
// @Param		end_date	 query		string                   	  false	"End Date"
// @Success		200	         {object}	models.RevenueByCategoryResp
// @Failure		400	         {object}	responsehelper.CommonResponse
// @Failure		500	         {object}	responsehelper.CommonResponse
// @Router		/api/v1/orders/revenue-by-category [get]
func (h *Handler) RevenueByCategory(c *gin.Context) {
	dateRange, sErr := urlquery.GetDateRangeQuery(c)
	if sErr != nil {
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	var resp models.RevenueByCategoryResp
	resp, sErr = h.svc.RevenueByCategory(dateRange)
	if sErr != nil {
		if sErr.Code >= http.StatusInternalServerError {
			h.logger.Error("Error fetching revenue by category", zap.Error(sErr.Error))
			c.JSON(sErr.Code, responsehelper.NewCommonResponse("Internal Server Error"))
			return
		}
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
