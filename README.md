# Golang API Service with Go-Clean-Architecture and Echo v4 Framework

This is a Golang API service built using the Go-Clean-Architecture design pattern and powered by Echo v4 framework. It comes with several features that are commonly used in API services.  The project includes various features such as database connection, ORM integration, database migration, REST API endpoints, default response handling, validation, payload binding, authentication, and middleware.


## Features

- ✅ Database Connection: The project establishes a connection to the database.
- ✅ Integrate ORM: Object-Relational Mapping (ORM) is integrated to simplify database interactions.
- ✅ Go Lang Migrate: Database migrations are managed using Go Lang Migrate.
- ✅ REST API: The API service exposes RESTful endpoints to interact with the application.
- ✅ Default Response: Default response handling is implemented to provide consistent API responses.
- ✅ Validation: Incoming API requests are validated to ensure data integrity.
- ✅ Binding Payload: API payloads are appropriately bound to data models for processing.
- ✅ Authentication: The API is secured with authentication mechanisms.
- ✅ Middleware: Custom middleware is implemented to handle specific API requests.

## Getting Started

### Prerequisites

Before running the application, make sure you have the following installed:

- Golang: [https://golang.org/](https://golang.org/)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/advengulo/go-clean-arch-test.git
cd go-clean-arch-test
```

2. Install dependencies:

```bash
go mod download
```

## Project Structure
```markdown
- cmd/
    - main.go               # Application entry point
- internal/
    - middlewares/          # Custom middleware
    - models/               # Data models
    - modules/
        - app/
            - handler.go     # API handlers
            - repository.go  # Database repository interfaces
            - usecase.go     # Business logic use cases
- pkg/
    - database/
        - migrations/       # Database migrations
    - utils/                # Utility functions
```

## Config
`.env`

```
APPLICATION_NAME = 

# [database]
DB_HOST = 
DB_PORT = 
DB_USER = 
DB_PASSWORD =
DB_NAME = 

# [config]
JWT_KEY_SECRET = 
```

Feel free to explore and extend the functionalities of this Golang API service as per your project requirements. Happy coding! 🚀