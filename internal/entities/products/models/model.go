package models

type Product struct {
	ID           string  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string  `gorm:"type:varchar(255);not null" json:"name"`
	Category     string  `gorm:"type:varchar(100);not null" json:"category"`
	UnitPrice    float64 `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	Discount     float64 `gorm:"type:decimal(10,2);default:0" json:"discount"`
	ShippingCost float64 `gorm:"type:decimal(10,2);default:0" json:"shipping_cost"`
}
