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

To run migrations

```
go run cmd/migrator/main.go
```

To run importer (csv importer and upload to db)

```
go run cmd/data-importer.go
```
