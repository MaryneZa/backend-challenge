# Backend-Challenge

## Project setup and run instructions

```
├───cmd
│   ├───grpc-server
│   └───rest-server
└───internal
    ├───adapter
    │   ├───config
    │   ├───grpc
    │   │   ├───proto
    │   │   └───stub
    │   ├───handler
    │   ├───middleware
    │   └───storage
    │       └───mongo
    │           └───repository
    └───core
        ├───domain
        ├───port
        ├───service
        └───util
            └───test
```

### Environment Setup

* Create `.env` file:

```bash
cp .example.env .env
```

### MongoDB Setup with Docker

```bash
docker-compose up -d
```

### Access MongoDB container:

```bash
docker exec -it <container-name> mongosh -u <user> -p <password>
```

### Run servers

* REST API Server:

```bash
go run cmd/rest-server/main.go
```

* gRPC API Server:

```bash
go run cmd/grpc-server/main.go
```

### Run tests

```bash
go test ./... -v
```

* Includes both unit and integration tests
* Use `.env` and `TEST_MONGO_...` vars for test DB

---

## JWT Token Usage Guide

All endpoints except `/register` and `/login` require JWT authentication.
Add this to the request header:

```
Authorization: Bearer <jwt-token>
```

---

## API Endpoints

### `POST /register`

**Request Body**

```json
{
  "email": "test@example.com",
  "password": "password"
}
```

**Response**

```json
{
  "token": "<token>"
}
```

---

### `POST /login`

**Request Body**

```json
{
  "email": "test@example.com",
  "password": "password"
}
```

**Response**

```json
{
  "token": "<token>"
}
```

---

### `GET /users`

**Response**

```json
[
  {
    "id": "...",
    "name": "",
    "email": "test@example.com",
    "password": "",
    "created_at": "..."
  }
]
```

---

### `GET /users/user/email`

**Request Body**

```json
{
  "email": "test@example.com"
}
```

**Response**

```json
{
  "id": "...",
  "name": "",
  "email": "test@example.com",
  "password": "",
  "created_at": "..."
}
```

---

### `GET /users/user/id`

**Request Body**

```json
{
  "id": "..."
}
```

**Response**

```json
{
  "id": "...",
  "name": "",
  "email": "test@example.com",
  "password": "",
  "created_at": "..."
}
```

---

### `PATCH /users/user/update-email`

**Request Body**

```json
{
  "email": "change@example.com"
}
```

**Response**

```json
{
  "message": "Update successfully!"
}
```

---

### `PATCH /users/user/update-name`

**Request Body**

```json
{
  "name": "new name"
}
```

**Response**

```json
{
  "message": "Update successfully!"
}
```

---

### `DELETE /users/user/delete`

**Request Body**

```json
{
  "email": "change@example.com"
}
```

**Response**

```json
{
  "message": "Delete change@example.com successfully!"
}
```

---
## gRPC Endpoints

####  Import the .proto file from `internal\adapter\grpc\proto\user.proto`

### CreateUser

**Request**

```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

**Response**

```json
{
  "message": "successful"
}
```

### GetUser

**Request**

```json
{
  "email": "stacy@example.com"
}
```

**Response**

```json
{
    "user": {
        "id": "...",
        "name": "",
        "email": "stacy@example.com",
        "created_at": "..."
    }
}
```


---

## References

### MongoDB

* [https://blog.me-idea.in.th/mongodb-docker-compose-up-%E0%B8%9B%E0%B8%B8%E0%B9%8A%E0%B8%9A%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87-mongo-database-%E0%B8%9B%E0%B8%B1%E0%B9%8A%E0%B8%9A-d27004a9fd78](https://blog.me-idea.in.th/mongodb-docker-compose-up-%E0%B8%9B%E0%B8%B8%E0%B9%8A%E0%B8%9A%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87-mongo-database-%E0%B8%9B%E0%B8%B1%E0%B9%8A%E0%B8%9A-d27004a9fd78)
* [https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
* [https://www.mongodb.com/docs/drivers/go/current/usage-examples](https://www.mongodb.com/docs/drivers/go/current/usage-examples)
* [https://www.mongodb.com/docs/drivers/go/current/usage-examples/count/#std-label-golang-count-usage-example](https://www.mongodb.com/docs/drivers/go/current/usage-examples/count/#std-label-golang-count-usage-example)

### Go Routine

* [https://docs.mikelopster.dev/c/goapi-essential/chapter-6/intro](https://docs.mikelopster.dev/c/goapi-essential/chapter-6/intro)

### Unit Test

* [https://docs.mikelopster.dev/c/goapi-essential/chapter-8/intro](https://docs.mikelopster.dev/c/goapi-essential/chapter-8/intro)

### Architecture (Hexagonal)

* [https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij](https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij)
* [https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
* [https://docs.mikelopster.dev/c/goapi-essential/chapter-7/intro](https://docs.mikelopster.dev/c/goapi-essential/chapter-7/intro)

### gRPC

* [https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05](https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05)
* [https://medium.com/@titlebhoomtawathplinsutมาทำ-grpc-service-ด้วย-go-กัน-866d7452f5dd](https://medium.com/@titlebhoomtawathplinsutมาทำ-grpc-service-ด้วย-go-กัน-866d7452f5dd)
