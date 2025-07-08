# Customer RESTful API

A clean and efficient RESTful API for customer management built with Go, following Clean Architecture principles and best practices.

## Features

- **Complete CRUD Operations**: Create, Read, Update, and Delete customers
- **Clean Architecture**: Organized with proper separation of concerns
- **Input Validation**: Comprehensive validation for all requests
- **Error Handling**: Centralized error handling with meaningful responses
- **Database Transactions**: Ensures data consistency with proper transaction management
- **CORS Support**: Cross-origin resource sharing enabled
- **Environment Configuration**: Configurable via environment variables

## Tech Stack

- **Language**: Go 1.19+
- **Database**: PostgreSQL
- **Router**: httprouter
- **Validation**: go-playground/validator
- **Environment**: godotenv

## Project Structure

```
customer-restful-api/
├── app/
│   └── database.go          # Database configuration
├── controller/
│   ├── customer_controller.go
│   └── customer_controller_impl.go
├── exception/
│   ├── error_handler.go     # Centralized error handling
│   └── not_found_error.go   # Custom error types
├── helper/
│   ├── cors.go              # CORS middleware
│   ├── error.go             # Error utilities
│   ├── json.go              # JSON utilities
│   ├── model.go             # Model converters
│   └── tx.go                # Transaction utilities
├── model/
│   ├── domain/
│   │   └── customer.go      # Domain models
│   └── web/
│       ├── customer_create_request.go
│       ├── customer_response.go
│       ├── customer_update_request.go
│       └── web_response.go
├── repository/
│   ├── customer_repository.go
│   └── customer_repository_impl.go
├── service/
│   ├── customer_service.go
│   └── customer_service_impl.go
├── .env                     # Environment variables
├── go.mod
├── go.sum
└── main.go
```

## Getting Started

### Prerequisites

- Go 1.19 or higher
- PostgreSQL 12 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/customer-restful-api.git
   cd customer-restful-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/customer_db?sslmode=disable
   PORT=8080
   ```

4. **Set up the database**
   
   Create a PostgreSQL database and table:
   ```sql
   CREATE DATABASE customer_db;
   
   \c customer_db;
   
   CREATE TABLE customers (
       id SERIAL PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       email VARCHAR(255) NOT NULL UNIQUE,
       phone VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## API Endpoints

### Base URL
```
http://localhost:8080/api
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/customers` | Get all customers |
| POST | `/customers` | Create a new customer |
| GET | `/customers/{id}` | Get customer by ID |
| PUT | `/customers/{id}` | Update customer by ID |
| DELETE | `/customers/{id}` | Delete customer by ID |

### Request/Response Examples

#### Create Customer
```bash
POST /api/customers
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "+1234567890"
}
```

**Response:**
```json
{
  "code": 201,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "+1234567890",
    "created_at": "2023-01-01T10:00:00Z",
    "updated_at": "2023-01-01T10:00:00Z"
  }
}
```

#### Get All Customers
```bash
GET /api/customers
```

**Response:**
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "+1234567890",
      "created_at": "2023-01-01T10:00:00Z",
      "updated_at": "2023-01-01T10:00:00Z"
    }
  ]
}
```

#### Get Customer by ID
```bash
GET /api/customers/1
```

#### Update Customer
```bash
PUT /api/customers/1
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "phone": "+1234567891"
}
```

#### Delete Customer
```bash
DELETE /api/customers/1
```

**Response:**
```json
{
  "code": 204,
  "status": "OK"
}
```

## Validation Rules

### Customer Create/Update
- **Name**: Required, 1-255 characters
- **Email**: Required, valid email format, max 255 characters
- **Phone**: Required, E.164 format, max 255 characters

### Error Response Format
```json
{
  "code": 400,
  "status": "BAD REQUEST",
  "data": [
    {
      "field": "email",
      "error": "This field must be a valid email"
    }
  ]
}
```

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o customer-api main.go
```

### Docker Support
```dockerfile
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o customer-api main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/customer-api .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./customer-api"]
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `PORT` | Server port | 8080 |

## Architecture

This project follows Clean Architecture principles:

- **Domain Layer**: Core business entities and rules
- **Repository Layer**: Data access abstraction
- **Service Layer**: Business logic and use cases
- **Controller Layer**: HTTP handlers and request/response logic
- **Helper Layer**: Utility functions and middleware

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Fathur Rohman - fr081938@gmail.com
Project Link: [https://github.com/FathurRohmna/customer-restful-api](https://github.com/FathurRohmna/customer-restful-api)