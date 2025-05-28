package handlers

import (
	"net/http"

	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/core/responsehelper"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	RevenueByProduct give stats of revenue by each product
//
// @Summary		stats of revenue by each product
// @Description	get revenue by product for a given date range
// @ID			revenue-by-product
// @Tags		revenue
// @Produce		json
// @Param		start_date	 query		string                   	  false	"Start Date"
// @Param		end_date	 query		string                   	  false	"End Date"
// @Success		200	         {object}	models.RevenueByProductResp
// @Failure		400	         {object}	responsehelper.CommonResponse
// @Failure		500	         {object}	responsehelper.CommonResponse
// @Router		/api/v1/orders/revenue-by-product [get]
func (h *Handler) RevenueByProduct(c *gin.Context) {
	dateRange, sErr := urlquery.GetDateRangeQuery(c)
	if sErr != nil {
		c.JSON(sErr.Code, responsehelper.NewCommonResponse(sErr.Error.Error()))
		return
	}
	var resp models.RevenueByProductResp
	resp, sErr = h.svc.RevenueByProduct(dateRange)
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
