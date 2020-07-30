# Order Analysis

## Features
- Converting huge CSV formatted file to multiple JSON files.
- gRPC based HTTP server to serve the results using the created JSON files.
- Swagger UI for the HTTP server at `/swagger-ui`.

## Usage

### Building protobuf files
```bash
make proto
```

### Processing data
```bash
make process-data
```

### Running the server
```bash
make run
```

### API's to be supported
## Authorization
1. `POST /login`
## Orders
1. `GET /orders/{id}`
2. `POST /orders`
3. `PUT /orders/{id}`
4. `DELETE /orders/{id}`
## Customers
1. `GET /customers`
2. `GET /customers/{id}`
3. `POST /customers/{id}`
4. `PUT /customers/{id}`
## Restaurants
1. `GET /restaurants/{id}`
2. `GET /restaurants/{name}`
3. `POST /restaurants`
4. `GET /restaurant/{id}/items/{item_id}`
5. `GET /restaurant/{id}/items?min={min}&max={max}`
6. `GET /restaurant/{id}/items`
7. `POST /restaurant/{id}/items`
8. `PUT /restaurant/{id}/items/{item_id}`
9. `DELETE /restaurant/{id}/items/{item_id}`
## Problem Statement
1. `GET /restaurants/top/{num}`
2. `GET /restaurants/worst/{num}`
3. `GET /state-cuisines/top/{num}`
4. `GET /state-cuisines/worst/{num}`
5. `GET /cuisines-demographics/{cuisine}`

