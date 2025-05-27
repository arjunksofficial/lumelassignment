package models

type SalesEntry struct {
	OrderID         int `csv:"Order ID"`
	ProductID       int
	CustomerID      int
	ProductName     string
	Category        string
	Region          string
	DateOfSale      string // Use appropriate type for date, e.g., time.Time
	QuantitySold    int
	UnitPrice       float64
	Discount        float64
	ShippingCost    float64
	PaymentMethod   string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
}
