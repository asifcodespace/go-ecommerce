# Go Ecommerce Server

A backend API server for an ecommerce application built with **Golang**.
This project follows a clean and modular structure with separate layers for domain, repository, services, database, and REST handlers.

## Features

* User registration and authentication
* JWT-based authentication
* Product management
* RESTful API endpoints
* PostgreSQL database integration
* Database migration support
* Repository and service layer separation
* Middleware support (CORS, Logger, Authentication)

## Tech Stack

* **Language:** Go
* **Database:** PostgreSQL
* **API:** REST API
* **Authentication:** JWT
* **Database Driver:** SQL
* **Migration:** SQL migration files

## Project Structure

```text
.
├── cmd/                 # Application commands
├── config/              # Configuration handling
├── db_queries/          # SQL queries
├── domain/              # Domain models
├── infra/
│   └── db/              # Database connection and migration
├── migrations/          # Database migration files
├── product/             # Product business logic
├── repo/                # Repository layer
├── rest/                # HTTP server, handlers, middleware
├── user/                # User business logic
├── util/                # Utility functions
├── go.mod
├── go.sum
└── main.go
```

## Requirements

Before running this project, make sure you have:

* Go installed
* PostgreSQL installed and running
* Required environment variables configured

## Environment Setup

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database

JWT_SECRET=your_secret_key
```

## Installation

Clone the repository:

```bash
git clone git@github.com:asifcodespace/go-ecommerce.git
```

Go to the project directory:

```bash
cd go-ecommerce
```

Install dependencies:

```bash
go mod download
```

## Database Setup

Create a PostgreSQL database and update your `.env` file.

Run migrations:

```bash
go run main.go
```

## Running the Application

Start the server:

```bash
go run main.go
```

The server will start on the configured port.

## API Overview

### User

* Register user
* Login user
* JWT authentication

### Product

* Create product
* Get products
* Get single product
* Update product
* Delete product

## Development

Format Go code:

```bash
gofmt -w .
```

Run tests:

```bash
go test ./...
```

## Environment Variables

Sensitive information such as:

* Database credentials
* JWT secrets
* API keys

should not be committed. Keep them inside `.env` and add them to `.gitignore`.

## License

This project is for learning and development purposes.
