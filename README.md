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

maybe need go get

```
go get
```

### run

```bash
go run .
```

## format

### when to must check

when open pull request

### format code

```bash
npm run format
```

### check code format

````bash
npm run lint
## sql

### start

> detach is background run, if want see its output remove --detach

```bash
cd sql
docker-compose up --detach
````

### into postgres

- for have psql user

```bash
psql -h localhost -p 5432 -U user
```

- for no psql user

```bash
docker exec -it <containerName> psql -U user
```

### put txt into sql

- for have psql user

```bash
psql -h localhost -p 5432 -U user < sql.txt
```

- for no psql user

```bash
docker exec -it <containerName> psql -U user < sql.txt
```

### close sql

```bash
docker-compose down
```
