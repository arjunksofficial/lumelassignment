package dataimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	customermodels "github.com/arjunksofficial/lumelassignment/internal/entities/customers/models"
	customerstore "github.com/arjunksofficial/lumelassignment/internal/entities/customers/store"
	ordermodels "github.com/arjunksofficial/lumelassignment/internal/entities/orders/models"
	orderstore "github.com/arjunksofficial/lumelassignment/internal/entities/orders/store"
	productmodels "github.com/arjunksofficial/lumelassignment/internal/entities/products/models"
	productstore "github.com/arjunksofficial/lumelassignment/internal/entities/products/store"
	"github.com/arjunksofficial/lumelassignment/pkg/data-importer/models"
	"github.com/jszwec/csvutil"
)

type Importer struct {
	ProductStore  productstore.Store
	CustomerStore customerstore.Store
	OrderStore    orderstore.Store
}

func NewImporter() *Importer {
	return &Importer{
		ProductStore:  productstore.GetStore(),
		CustomerStore: customerstore.GetStore(),
		OrderStore:    orderstore.GetStore(),
	}
}

func (i *Importer) ProcessSalesData() error {
	fileName := "sales_data.csv"
	fileReader, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open CSV file %s: %w", fileName, err)
	}
	defer fileReader.Close()
	csvReader := csv.NewReader(fileReader)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	var salesEntries []models.SalesEntry
	for {
		var entry models.SalesEntry
		if err := dec.Decode(&entry); err == io.EOF {
			break // End of file
		} else if err != nil {
			return fmt.Errorf("failed to decode CSV entry: %w", err)
		}
		salesEntries = append(salesEntries, entry)
	}
	if len(salesEntries) == 0 {
		return fmt.Errorf("no sales entries found in the CSV file %s", fileName)
	}
	if err := i.processSalesEntries(salesEntries); err != nil {
		return fmt.Errorf("failed to process sales entries: %w", err)
	}
	log.Printf("Successfully processed %d sales entries from %s", len(salesEntries), fileName)
	return nil
}

func (i *Importer) processSalesEntries(entry []models.SalesEntry) error {
	// Here you can implement the logic to process each sales entry
	// upload products and customers if not exists

	products := []productmodels.Product{}
	customers := []customermodels.Customer{}
	orders := []ordermodels.Order{}
	productsMap := make(map[string]productmodels.Product)    // To avoid duplicates
	customersMap := make(map[string]customermodels.Customer) // To avoid duplicates
	// Iterate over each sales entry
	for _, sale := range entry {

		// Process product
		product := productmodels.Product{
			ID:           sale.ProductID,
			Name:         sale.ProductName,
			Category:     sale.Category,
			UnitPrice:    sale.UnitPrice,
			Discount:     sale.Discount,
			ShippingCost: sale.ShippingCost,
		}
		// Check if product already exists in the map
		if existingProduct, exists := productsMap[product.ID]; exists {
			// If it exists, update the existing product
			existingProduct.Name = product.Name
			existingProduct.Category = product.Category
			existingProduct.UnitPrice = product.UnitPrice
			existingProduct.Discount = product.Discount
			existingProduct.ShippingCost = product.ShippingCost
			productsMap[product.ID] = existingProduct // Update the map with the modified product
		} else {
			// If it doesn't exist, add the new product to the map
			productsMap[product.ID] = product
		}

		// Process customer
		customer := customermodels.Customer{
			ID:      sale.CustomerID, // Assuming CustomerID is unique
			Name:    sale.CustomerName,
			Email:   sale.CustomerEmail,
			Address: sale.CustomerAddress,
		}

		// Check if customer already exists in the map
		if existingCustomer, exists := customersMap[customer.ID]; exists {
			// If it exists, update the existing customer
			existingCustomer.Name = customer.Name
			existingCustomer.Email = customer.Email
			existingCustomer.Address = customer.Address
			customersMap[customer.ID] = existingCustomer // Update the map with the modified customer
		} else {
			// If it doesn't exist, add the new customer to the map
			customersMap[customer.ID] = customer
		}

		// Convert date of sale from string to time.Time
		dateOfSale, err := time.Parse("2006-01-02", sale.DateOfSale)
		if err != nil {
			return fmt.Errorf("failed to parse date of sale %s: %w", sale.DateOfSale, err)
		}
		// Process order
		order := ordermodels.Order{
			ID:            sale.OrderID,
			ProductID:     product.ID,
			CustomerID:    customer.ID,
			DateOfSale:    dateOfSale,
			QuantitySold:  sale.QuantitySold,
			PaymentMethod: sale.PaymentMethod,
			Region:        sale.Region,
		}
		orders = append(orders, order)
	}
	// Convert maps to slices
	for _, product := range productsMap {
		products = append(products, product)
	}
	for _, customer := range customersMap {
		fmt.Println(customer.ID)
		customers = append(customers, customer)
	}

	err := i.ProductStore.BulkCreateOrUpdate(products)
	if err != nil {
		return fmt.Errorf("failed to bulk create or update products: %w", err)
	}
	err = i.CustomerStore.BulkCreateOrUpdate(customers)
	if err != nil {
		return fmt.Errorf("failed to bulk create or update customers: %w", err)
	}
	err = i.OrderStore.BulkCreateOrUpdate(orders)
	if err != nil {
		return fmt.Errorf("failed to bulk create or update orders: %w", err)
	}
	log.Printf("Successfully processed %d products, %d customers, and %d orders", len(products), len(customers), len(orders))

	return nil
}
