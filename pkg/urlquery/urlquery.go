package urlquery

import (
	"errors"
	"net/http"
	"time"

	"github.com/arjunksofficial/lumelassignment/pkg/core/serror"
	"github.com/gin-gonic/gin"
)

type DateRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func GetDateRangeQuery(c *gin.Context) (DateRange, *serror.ServiceError) {
	dateRange := DateRange{}
	currentDate := time.Now()
	if c.Query("start_date") != "" {
		startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
		if err != nil {
			return dateRange, &serror.ServiceError{
				Code:  http.StatusBadRequest,
				Error: errors.New("Invalid start date format"),
			}
		}
		dateRange.StartDate = startDate
	} else {
		dateRange.StartDate = currentDate.AddDate(0, 0, -30) // Default to 30 days ago
	}
	if c.Query("end_date") != "" {
		endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
		if err != nil {
			return dateRange, &serror.ServiceError{
				Code:  http.StatusBadRequest,
				Error: errors.New("Invalid end date format"),
			}
		}
		dateRange.EndDate = endDate
	} else {
		dateRange.EndDate = currentDate // Default to today
	}
	if dateRange.StartDate.After(dateRange.EndDate) {
		return dateRange, &serror.ServiceError{
			Code:  http.StatusBadRequest,
			Error: errors.New("Start date cannot be after end date"),
		}
	}
	return dateRange, nil

}
