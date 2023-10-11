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

not yet use

```bash

```

### run

```bash
go run .
```

## sql

### start

> detach is background run, if want see its output remove --detach

```bash
cd sql
docker-compose up --detach
```

### into postgres

```bash
psql -h localhost -p 5432 -U user
```

### close sql

```bash
docker-compose down
```

### put txt into sql

```bash
psql -h localhost -p 5432 -U user < sql.txt
```
