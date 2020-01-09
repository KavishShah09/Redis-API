# Simple GO Lang REST API

> Simple RESTful API to create and read Redis database "key" with "values"

## Quick Start

### Prerequisites

```
Redis must be installed and a redis server must be running on PORT:6379 (i.e. default port of redis database)
```

```bash
# Clone the repository
git clone https://github.com/KavishShah09/Redis-API.git
```

```bash
# Install mux router and go-redis
go get -u github.com/gorilla/mux
go get -u github.com/go-redis/redis
```

```bash
go run main.go
# Server will start running
```

## Endpoints

### Get All Data

```
GET request localhost:8000/data  // with query params


# Sample Request With Query Params (default limit = 10 and match = "")
# localhost:8000/data?limit=10&match=Ka*
# Will get 10 values starting with the letters "Ka"
```

### Get Single Key Value

```
GET request localhost:8000/data/{key}
```

### Create Key and Value

```
POST request localhost:8000/data

# Request sample
# {
#   "key":"firstname",
#   "value":"Kavish",
# }
```

## App Info

### Author

[Kavish Shah](http://www.linkedin.com/in/kavish-shah-501b32192)

### Version

1.0.0
