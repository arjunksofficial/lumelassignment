//	@title			Lumel Assignment
//	@version		1.0
//	@description	This is a sample server for serving revenue stats

//	@contact.name	API Support

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@BasePath	/api/v1

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
package main

import (
	"log"

	"github.com/arjunksofficial/lumelassignment/pkg/api"
	"github.com/arjunksofficial/lumelassignment/pkg/config"
)

func main() {
	log.Println("Starting API server...")
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}
	log.Printf("Configuration loaded: %+v", config)
	r := api.GetRouter()
	if err := r.Run(":" + config.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
