# lumelassignment

## Initial Set up

Run the following to start postgres in docker

```
cd build/docker
docker-compose up -d
cd ..

```

To setup postgres database after running migrations to create tables, indices, etc

```
source config/local/.env
go run cmd/migrator/main.go
```

To run importer (csv importer and upload to db) to populate the database with data

```
go run cmd/data-importer.go
```

To Run api-server which servces revenue stats endpoint

```
go run cmd/apiserver/main.go
```

Refer docs/swagger/swagger.yaml for API documentation
To run tests

```
go test ./...
```
