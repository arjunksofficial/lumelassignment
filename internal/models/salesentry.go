package models

type SalesEntry struct {
	OrderID         int    `csv:"Order ID"`
	ProductID       string `csv:"Product ID"`
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

// 	- **Order ID:** (Unique identifier)
// - **Product ID:** (Unique identifier)
// - **Customer ID:** (Unique identifier)
// - **Region:**
// - **Date of Sale:**
// - **Quantity Sold:**
// - **Shipping Cost:**
// - **Payment Method:**
