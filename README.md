# Project golang-starter

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.


## Project Structure Analysis

### Visual Representation
```
.
├── config/
│   └── config.go
├── internal/
│   └── entrypoint/
│       ├── http/
│       │   ├── dtos.go
│       │   ├── middleware/
│       │   ├── router/
│       │   └── user/
├── pkg/
│   ├── aws/
│   ├── database/
│   ├── http.router.middlewares/
│   ├── logger/
│   └── server/
├── vendor/
│   └── (third-party dependencies)
├── docker-compose.yml
├── .goreleaser.yml
└── README.md
```

### Folder Structure Description

1. **Root Level**
    - `docker-compose.yml`: Contains Docker container configuration
    - `.goreleaser.yml`: Configuration for building and releasing Go applications
    - `README.md`: Project documentation

2. **config/**
    - Contains configuration related code
    - Manages application configuration settings

3. **internal/**
    - Contains private application code
    - Not importable by other projects
    - Structure:
        - `entrypoint/`: Application entry points
        - `http/`: HTTP related implementations
        - `middleware/`: Custom HTTP middleware
        - `router/`: Routing configuration
        - `user/`: User-related handlers

4. **pkg/**
    - Contains public packages that can be imported by other projects
    - Components:
        - `aws/`: AWS integration code
        - `database/`: Database related code
        - `http.router.middlewares/`: HTTP router middleware
        - `logger/`: Logging functionality
        - `server/`: Server implementation

5. **vendor/**
    - Contains third-party dependencies
    - Managed by Go modules
    - Includes various external libraries:
        - AWS SDK
        - Gin Web Framework
        - PostgreSQL drivers
        - Other utility libraries

### Responsibilities

1. **Configuration Management**
    - The `config` package manages application settings and environment variables
    - Provides centralized configuration handling

2. **HTTP Layer**
    - Located in `internal/entrypoint/http`
    - Handles HTTP requests and responses
    - Implements routing and middleware
    - Manages DTOs (Data Transfer Objects)

3. **Infrastructure**
    - AWS integration (`pkg/aws`)
    - Database operations (`pkg/database`)
    - Logging functionality (`pkg/logger`)
    - Server configuration (`pkg/server`)

4. **Middleware**
    - Request/response processing
    - Authentication and authorization
    - Logging and monitoring
    - Error handling

5. **Business Logic**
    - User management (`internal/entrypoint/http/user`)
    - Service layer implementations
    - Domain logic and operations

## Design Patterns

1. **Clean Architecture**
    - Separation of concerns
    - Clear dependency boundaries
    - Internal vs. public packages

2. **Dependency Injection**
    - Modular design
    - Testable components
    - Loose coupling

3. **Middleware Pattern**
    - Request/response pipeline
    - Chainable middleware components
    - Cross-cutting concerns

4. **Repository Pattern**
    - Database abstraction
    - Data access layer
    - Separation from business logic

This structure follows Go best practices and conventions, providing a clean, maintainable, and scalable architecture for the application.


## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```