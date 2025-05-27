package dataimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	customerstore "github.com/arjunksofficial/lumelassignment/internal/entities/customers/store"
	orderstore "github.com/arjunksofficial/lumelassignment/internal/entities/orders/store"
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

	header := dec.Header()
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
	// read using
	fmt.Println("CSV Header:", header)
	for _, entry := range salesEntries {
		// Process each sales entry
		if err := i.processSalesEntry(entry); err != nil {
			return fmt.Errorf("failed to process sales entry: %w", err)
		}
	}

	return nil
}

func (i *Importer) processSalesEntry(entry models.SalesEntry) error {
	fmt.Println("Processing Sales Entry:", entry)
	// Here you can implement the logic to process each sales entry
	// upload products and customers if not exists
	return nil
}
