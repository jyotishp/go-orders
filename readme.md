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