package store

import (
	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/database"
	"github.com/arjunksofficial/lumelassignment/pkg/urlquery"
	"gorm.io/gorm"
)

type Store interface {
	BulkCreate(customers []models.Order) error
	BulkCreateOrUpdate(customers []models.Order) error

	RevenueTotal(filter urlquery.DateRange) (float64, error)
	RevenueByProduct(filter urlquery.DateRange) ([]models.ProductRevenue, error)
	RevenueByCategory(filter urlquery.DateRange) ([]models.CategoryRevenue, error)
	RevenueByRegion(filter urlquery.DateRange) ([]models.RegionRevenue, error)
	RevenueByTrends(filter urlquery.DateRange) ([]models.TrendRevenue, error)
}

type store struct {
	db *gorm.DB
}

func GetStore() Store {
	return &store{
		db: database.GetPostgresDB(),
	}
}

func (s *store) BulkCreate(customers []models.Order) error {
	if len(customers) == 0 {
		return nil // No customers to create
	}

	// Use the Create method to insert multiple records
	result := s.db.Create(&customers)
	if result.Error != nil {
		return result.Error // Return any error that occurred during the creation
	}

	return nil // Return nil if the operation was successful
}

func (s *store) BulkCreateOrUpdate(customers []models.Order) error {
	if len(customers) == 0 {
		return nil // No customers to create or update
	}

	// Use the Save method to insert or update multiple records
	result := s.db.Save(&customers)
	if result.Error != nil {
		return result.Error // Return any error that occurred during the save operation
	}

	return nil // Return nil if the operation was successful
}

func (s *store) RevenueTotal(filter urlquery.DateRange) (float64, error) {
	var total float64
	query := s.db.Model(&models.Order{}).Joins("JOIN products ON orders.product_id = products.id").
		Select("SUM((quantity_sold * products.unit_price - discount) + products.shipping_cost) AS total_revenue")
	if filter.StartDate.IsZero() && filter.EndDate.IsZero() {
		query = query.Where("date_of_sale BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	} else if !filter.StartDate.IsZero() {
		query = query.Where("date_of_sale >= ?", filter.StartDate)
	} else if !filter.EndDate.IsZero() {
		query = query.Where("date_of_sale <= ?", filter.EndDate)
	}
	err := query.Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *store) RevenueByProduct(filter urlquery.DateRange) ([]models.ProductRevenue, error) {
	var revenues []models.ProductRevenue
	query := s.db.Model(&models.Order{}).
		Joins("JOIN products ON orders.product_id = products.id").
		Select("products.id AS product_id, products.name AS product_name, SUM((orders.quantity_sold * (products.unit_price - products.discount + products.shipping_cost))) AS revenue").
		Group("products.id")
	if filter.StartDate.IsZero() && filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	} else if !filter.StartDate.IsZero() {
		query = query.Where("orders.date_of_sale >= ?", filter.StartDate)
	} else if !filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale <= ?", filter.EndDate)
	}
	err := query.Scan(&revenues).Error
	if err != nil {
		return nil, err
	}
	if len(revenues) == 0 {
		return nil, nil // No revenues found
	}

	return revenues, nil
}
func (s *store) RevenueByCategory(filter urlquery.DateRange) ([]models.CategoryRevenue, error) {
	var revenues []models.CategoryRevenue
	query := s.db.Model(&models.Order{}).
		Joins("JOIN products ON orders.product_id = products.id").
		Select("products.category AS category, SUM((orders.quantity_sold * (products.unit_price - products.discount + products.shipping_cost))) AS revenue").
		Group("products.category")
	if filter.StartDate.IsZero() && filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}
	if !filter.StartDate.IsZero() {
		query = query.Where("orders.date_of_sale >= ?", filter.StartDate)
	}
	if !filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale <= ?", filter.EndDate)
	}
	err := query.Scan(&revenues).Error
	if err != nil {
		return nil, err
	}
	if len(revenues) == 0 {
		return nil, nil // No revenues found
	}
	return revenues, nil
}
func (s *store) RevenueByRegion(filter urlquery.DateRange) ([]models.RegionRevenue, error) {
	var revenues []models.RegionRevenue
	query := s.db.Model(&models.Order{}).
		Joins("JOIN products ON orders.product_id = products.id").
		Select("products.region AS region, SUM((orders.quantity_sold * (products.unit_price - products.discount + products.shipping_cost))) AS revenue").
		Group("products.region")
	if filter.StartDate.IsZero() && filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}
	if !filter.StartDate.IsZero() {
		query = query.Where("orders.date_of_sale >= ?", filter.StartDate)
	}
	if !filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale <= ?", filter.EndDate)
	}
	err := query.Scan(&revenues).Error
	if err != nil {
		return nil, err
	}
	if len(revenues) == 0 {
		return nil, nil // No revenues found
	}
	return revenues, nil
}

func (s *store) RevenueByTrends(filter urlquery.DateRange) ([]models.TrendRevenue, error) {
	var revenues []models.TrendRevenue
	query := s.db.Model(&models.Order{}).
		Joins("JOIN products ON orders.product_id = products.id").
		Select("DATE_FORMAT(orders.date_of_sale, '%Y-%m') AS month, SUM((orders.quantity_sold * (products.unit_price - products.discount + products.shipping_cost))) AS revenue").
		Group("DATE_FORMAT(orders.date_of_sale, '%Y-%m')")
	if filter.StartDate.IsZero() && filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}
	if !filter.StartDate.IsZero() {
		query = query.Where("orders.date_of_sale >= ?", filter.StartDate)
	}
	if !filter.EndDate.IsZero() {
		query = query.Where("orders.date_of_sale <= ?", filter.EndDate)
	}
	err := query.Scan(&revenues).Error
	if err != nil {
		return nil, err
	}
	if len(revenues) == 0 {
		return nil, nil // No revenues found
	}
	return revenues, nil
}
