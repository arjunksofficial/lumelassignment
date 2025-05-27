package main

import (
	"fmt"

	customermmodels "github.com/arjunksofficial/lumelassignment/internal/entities/customers/models"
	ordermodels "github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	productmodels "github.com/arjunksofficial/lumelassignment/internal/entities/products/models"

	"github.com/arjunksofficial/lumelassignment/pkg/core/logger"

	"github.com/arjunksofficial/lumelassignment/pkg/database"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("Connecting to Postgres database...")
	db := database.GetPostgresDB()
	logger.Info("Connected to Postgres database successfully.")
	logger.Info("Starting migration...")
	// 🔧 Auto-migrate your models
	err := db.AutoMigrate(&productmodels.Product{}, &customermmodels.Customer{}, &ordermodels.Order{})
	if err != nil {
		panic("migration failed: " + err.Error())
	}
	fmt.Println("Migration completed.")
}
