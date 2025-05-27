package models

type Customer struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Email   string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Address string `gorm:"type:varchar(255);not null" json:"address"`
}
