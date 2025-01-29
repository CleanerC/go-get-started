# Go-Kit HTTP Microservice Example

A basic HTTP microservice implementation using Go-Kit that performs addition of two numbers.

## Project Structure

```
.
├── main.go        # Entry point of the application 
├── go.mod         # Go module definition 
├── go.sum         # Go module checksums 
└── pkg/ 
    └── v1/ 
        ├── endpoint/   # Contains endpoint definitions 
        ├── service/    # Business logic implementation 
        └── transport/  # HTTP transport layer
```

## Components

### Service Layer (`pkg/v1/service/service.go`)
- Defines the core business logic interface
- Implements addition operation
- Contains data structures for requests/responses

### Endpoint Layer (`pkg/v1/endpoint/endpoint.go`)
- Defines endpoint adapters
- Converts transport-domain requests to service-domain requests
- Handles endpoint routing

### Transport Layer (`pkg/v1/transport/http.go`)
- HTTP transport implementation
- Request/response encoding/decoding
- Error handling
- Route definitions

## Setup

```bash
go mod tidy
go run main.go
```

## Usage

Send a GET request to add two numbers:
```bash
curl "http://localhost:8888/add?a=5&b=3"
```

### Expected Response:
```json
{"res": 8}
```

## API Endpoints

### GET /add?a={number}&b={number}
Adds two numbers

**Parameters:**
- `a`: First number
- `b`: Second number

**Returns:** JSON object with result `{"res": sum}`
