package main

import (
	dataimporter "github.com/arjunksofficial/lumelassignment/pkg/data-importer"
)

func main() {
	importer := dataimporter.NewImporter()
	if err := importer.ProcessSalesData(); err != nil {
		panic("failed to process sales data: " + err.Error())
	}
	println("Sales data processed successfully.")
}
