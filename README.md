# lumelassignment

Set up
Run the following to start postgres in docker

```
cd build/docker
docker-compose up -d
cd ..

```

To setup postgres

```
source config/local/.env
```

To run migrations to create tables, indices, etc

```
go run cmd/migrator/main.go
```

To run importer (csv importer and upload to db) to populate the database with data

```
go run cmd/data-importer.go
```
