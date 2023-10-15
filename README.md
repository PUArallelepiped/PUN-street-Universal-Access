# PUN-street-access

PUA

# how to run it

## frontend

### env

```
cd frontend
npm install
```

### run

```bash
cd frontend
npm run dev -- --open
```

it will auto reload after code change

## backend

### env

```
go mod download
```

### run

```bash
go run cmd/main.go
```

## sql

### start

> detach is background run, if want see its output remove --detach

```bash
docker-compose up --detach
```

### into postgres

- for have psql user

```bash
psql -h localhost -p 5432 -U postgres
```

- for no psql user

```bash
docker exec -it <containerName> psql -U postgres
```

### close sql

```bash
docker-compose down
```
