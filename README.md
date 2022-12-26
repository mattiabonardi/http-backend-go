# Http Backend GO
**Http backend GO** is a generic backend template based on [Gin framework](https://github.com/gin-gonic/gin). The application has already integrated:

- Monitoring API (livez, readyz, metrics)
- Authentication API based on [JWT](https://github.com/golang-jwt/jwt)
- Swagger docs

## Installation

```go
go get -d -v
```

## Startup

```go
go run .
```

Api docs: http://localhost:8080/public/swagger/index.html

## Docker

Build and run

```go
sh build.sh
sh start.sh
```