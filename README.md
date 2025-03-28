# Company API

A REST API for managing companies and their owners, built with Go and CockroachDB.

## Features

- **Company Management**: Create, update, retrieve, and list companies.
- **Owner Management**: Add owners to companies, including name and SSN.
- **SSN Validation**: Simulate SSN validation with role-based access control.
- **Modular Design**: Supports multiple commands (e.g., HTTP and gRPC servers).

## Requirements

- Go 1.18+ installed locally.
- Podman for containerized deployment.
- CockroachDB instance for database operations.

## Project Structure

```
cmd/
  serve_http/    # HTTP server entry point
  serve_grpc/    # (Optional) gRPC server entry point
config/          # Configuration files and modules
internal/
  db/            # Database connection and setup
  models/        # Data models
  repositories/  # Database repositories
  server/        # HTTP server and handlers
  services/      # Business logic and utilities
pkg/
  logger/        # Logging utilities
migrations/      # Database migration scripts
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/froppa/company-api.git
   cd company-api
   ```

2. Run the application using the provided `run.sh` script:

   ```bash
   ./run.sh
   ```

   > **Note**: `docker-compose` was not used due to configuration issues. The `run.sh` script handles starting the services using `podman` with a sleep.. Might be my Mac + Docker that is being annoying.

## API Endpoints

### Companies

- **POST /companies**: Create a new company.
- **GET /companies**: List all companies.
- **GET /companies/{id}**: Retrieve a company by ID.
- **PUT /companies/{id}**: Update a company by ID.

### Owners

- **POST /companies/{id}/owners**: Add an owner to a company.

### SSN Validation

- **POST /auth/verify**: Validate an SSN with role-based access.

## Configuration

The application uses a `base.yml` file for configuration. Example:

```yml
server:
  port: 8081

database:
  host: localhost
  port: 26257
  user: root
  name: companyapi
  sslmode: disable
```

## Testing

Run the provided test script:

```bash
./test.sh
```
