package models

type Order struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID    int    `gorm:"not null" json:"product_id"`
	CustomerID   int    `gorm:"not null" json:"customer_id"`
	Region       string `gorm:"type:varchar(100);not null" json:"region"`
	DateOfSale   string `gorm:"type:date;not null" json:"date_of_sale"` // Use appropriate type for date, e.g., time.Time
	QuantitySold int    `gorm:"not null" json:"quantity_sold"`
}
