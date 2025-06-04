# BUG TRACKER BACKEND

Instructions on how to manage the backend of the project is explained below. The instructions presume that you are already in the _bug-tracker-backend_ folder when you're running commands. So, before doing anything, naviagate to that folder by running `cd bug-tracker-backend` on your terminal.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Tools And Technologies](#tools-and-technologies)
- [Setting Up The Project](#setting-up-the-project)
    - [Creating .env](#creating-env)
    - [Creating The Database](#creating-the-database)
- [Database](#database)
    - [Creating The Tables](#creating-the-tables)
- [Running The Server](#running-the-server)
    - [Using Go](#using-go)
    - [Using Compile Daemon](#using-compile-daemon)
- [Project Conventions](#project-conventions)
    - [Naming Conventions](#naming-conventions)
    - [Creating Files](#creating-files)
    - [Handling Responses](#handling-responses)

## Prerequisites
- Go >= 1.24.3
- Postgres >= 16.0

## Tools And Technologies

- **Programming Language:** Go
- **Backend Framework:** Gin
- **ORM:** GORM
- **SQL Database:** Postgres
- **NoSQL Database:** Elascticsearch

## Setting Up The Project

### Creating .env

Create a `.env` file on the project root (inside the bug-tracker-backend) by copying the contents of the `.env.example` file and filling in the variables with the appropriate values.

### Creating The Database

To create the database from your terminal, activate `psql` with default user `postgres` using the command -

```bash
psql -U postgres
```

This should activate the `psql` CLI. Once you're in there, run the following SQL to create the database -

```sql
CREATE DATABASE bug_tracker_db;
```

To leave the `psql` CLI, just run `\q`.

## Database

### Creating The Tables

Run the following command on your terminal

```bash
go run migrations/migrate.go
```

## Running The Server

### Using Go

Run the following command on your terminal

```bash
go run main.go
```

To stop the server, press `Ctrl+C`.

### Using Compile Daemon

1. Install Compile Daemon by running the following command on your terminal -

    ```bash
    go install github.com/githubnemo/CompileDaemon@latest
    ```

2. Run the project using the following command

    ```bash
    CompileDaemon --command="./bug-tracker-backend"
    ```

    To stop the server, press `Ctrl+C`.

## Project Conventions

### Naming Conventions

- Use **Snake Case** _(file_name.go)_ to name files.
- Use **Camel Case** _(variableName)_ to name private variables and **Pascal Case** _(VariableName)_ to name public variables.
- Package names should be in lowercase and in one word.

### Creating Files

### Handling Responses

```go
// Handler that uses regular gin context (will be wrapped by middleware)
func regularHandler(c *gin.Context) {
	data := map[string]interface{}{
		"id":   "123",
		"name": "John Doe",
	}
	c.JSON(http.StatusOK, data)
}
```

```go
// Handler that uses enhanced context (bypasses middleware wrapping)
func enhancedHandler(c *gin.Context) {
	ec := GetEnhancedContext(c)
	
	data := map[string]interface{}{
		"id":   "456",
		"name": "Jane Doe",
	}
	
	ec.Success("User retrieved successfully", data)
}
```