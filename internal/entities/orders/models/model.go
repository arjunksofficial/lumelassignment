package models

import "time"

type Order struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     string    `gorm:"not null" json:"product_id"`
	CustomerID    string    `gorm:"not null" json:"customer_id"`
	Region        string    `gorm:"type:varchar(100);not null" json:"region"`
	DateOfSale    time.Time `gorm:"type:date;not null" json:"date_of_sale"`
	QuantitySold  int       `gorm:"not null" json:"quantity_sold"`
	PaymentMethod string    `gorm:"type:varchar(50);not null" json:"payment_method"`
	CreatedAt     time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp" json:"updated_at"`
}
