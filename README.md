# Rest
Logo

![ci](https://github.com/rest-go/rest/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/rest-go/rest/branch/main/graph/badge.svg?token=T38FWXMVY1)](https://codecov.io/gh/rest-go/rest)

Rest serves a fully RESTful API from any (PostgreSQL/MySQL/SQLite) database.

Visit https://rest-go.com for the full documentation, examples and guides.

## Install

There are various ways of installing Rest.

#### Precompiled binaries
Precompiled binaries for released versions are available in the [Releases page](https://github.com/rest-go/rest/releases). Using the latest production release binary is the recommended way of installing Rest. See the [INSTALLATION]() chapter in the documentation for all the details.

#### Go install

``` bash
go install github.com/rest-go/rest
```

## Run rest server
``` bash
# PG
rest -db.url "postgres://user:passwd@localhost:5432/db?search_path=api"

# MySQL
rest -db.url "mysql://user:passwd@tcp(localhost:3306)/db"

# SQLite
rest -db.url "sqlite://chinook.db"
```

## Use API

``` bash
# Create an artist
curl -XPOST "localhost:3000/artists" -d '{"artistid":10000, "name": "Bruce Lee"}'

# Read an artist
curl -XGET "localhost:3000/artists?&artistid=eq.10000"

# Update
curl -XPUT "localhost:3000/artists?&artistid=eq.10000" -d '{"name": "Stephen Chow"}'

# Delete
curl -XDELETE "localhost:3000/artists?&artistid=eq.10000"
```

## Docker image

``` bash
# for mysql
docker run -p 3000:3000 restgo/rest -db.url "mysql://user:passwd@tcp(host:port)/db"

# for sqlite with mounted volume
docker run -p 3000:3000 -v $(pwd):/data restgo/rest -db.url "sqlite:///data/chinook.db"
```

## JSON

``` bash
# POST json
curl -XPOST "localhost:3000/people" -d '{"id":1, "json_data": {"blood_type":"A-", "phones":[{"country_code":61, "number":"919-929-5745"}]}}'

# Fetch json field
curl "http://localhost:3000/people?select=id,json_data->>blood_type,json_data->>phones"
```

## Use rest as a Go library
It also works to embed rest server into an existing Go http server

``` go
package main

import (
	"log"
	"net/http"

	"github.com/rest-go/rest/pkg/server"
)

func main() {
	s := server.NewServer("sqlite://chinook.db")
	http.Handle("/", s)
	// or with prefix
	// http.Handle("/admin", s.WithPrefix("/admin"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Features
- [x] CRUD
- [x] Page
- [x] common operators e.g. eq,lt, gt, is, like
- [x] common types(int, bool, char, timestamp, decimal)
- [x] select fileds
- [x] order by
- [x] count
- [x] debug output sql & args
- [x] test
- [ ] security sql
- [ ] auth(http & jwt)
- [ ] comment/documentation
	- [ ] json
- [x] json (postgres, operations, nested post/get)
  - [x] quote
- [x] json (mysql & sqlite)
- [x] test for different db (github action)
- [ ] dump
## Road map
- [ ] Resource Embedding(one,many)
- [ ] open api
- [ ] Logical operators(or, and is already in code)
- [ ] escape field name
- [ ] application/x-www-form-urlencoded
- [ ] web management
