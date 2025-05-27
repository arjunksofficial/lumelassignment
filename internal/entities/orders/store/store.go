package store

import (
	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	"github.com/arjunksofficial/lumelassignment/pkg/database"
	"gorm.io/gorm"
)

type Store interface {
	BulkCreate(customers []models.Order) error
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
