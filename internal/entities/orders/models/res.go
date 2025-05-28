package models

type RevenueTotalResp struct {
	TotalRevenue float64 `json:"total_revenue"`
}

type RevenueByProductResp struct {
	Products []ProductRevenue `json:"products"`
}
type ProductRevenue struct {
	ProductID   string  `json:"product_id"`
	Revenue     float64 `json:"revenue"`
	ProductName string  `json:"product_name"`
}

type RevenueByRegionResp struct {
	Regions []RegionRevenue `json:"regions"`
}
type RegionRevenue struct {
	Region  string  `json:"region"`
	Revenue float64 `json:"revenue"`
}
type RevenueByCategoryResp struct {
	Categories []CategoryRevenue `json:"categories"`
}
type CategoryRevenue struct {
	Category string  `json:"category"`
	Revenue  float64 `json:"revenue"`
}
type RevenueByTrendsResp struct {
	Trends []TrendRevenue `json:"trends"`
}
type TrendRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}
