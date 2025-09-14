# BUG TRACKER BACKEND

Instructions on how to manage the backend of the project are given below. The instructions presume that you are already in the _bug-tracker-backend_ folder when you're running commands. So, before doing anything, naviagate to that folder by running `cd bug-tracker-backend` on your terminal.

[Apidog Documentation](https://build-together.apidog.io)

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
    - [Extracting Objects From The Request Context](#extracting-objects-from-the-request-context)

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
go run server/main.go
```

To stop the server, press `Ctrl+C`.

### Using Compile Daemon

1. Install Compile Daemon by running the following command on your terminal -

    ```bash
    go install github.com/githubnemo/CompileDaemon@latest
    ```

2. Run the project using the following command

    ```bash
    CompileDaemon -exclude-dir=api -build="go build -o bug-tracker-backend ./server" -command="./bug-tracker-backend"
    ```

    To stop the server, press `Ctrl+C`.

## Project Conventions

### Naming Conventions

- Use **Snake Case** _(file_name.go)_ to name files.
- Use **Camel Case** _(variableName)_ to name private variables and **Pascal Case** _(VariableName)_ to name public variables.
- Package names should be in lowercase and in one word.

### Creating Files

This project follows _MVC_ structure for organisation. Therefore, depending on the context of the code, they should be written in the appropriate module.

- **Routes:** Contains all the API routes/paths.
- **Controllers:** Contains all the API contollers, that is, the handler functions.
- **Models:** Contains the database models.
- **Migrations:** Contains the database migration files.
- **Middlewares:** Contains the system middlewares.
- **Types:** Contains the API request and response schemas.
- **Conf:** Contains the system configurations.
- **Utils:** Contains the system utility functions.

### Handling Responses

The [response middlewares](internal/middlewares/response_middleware.go) are responsible for setting a global standard response object for every API on the system. They wrap regular `json` responses with the global standard structure.

The following example illustrates how it can be ensured that the response object is properly set -

```go
// Handler that uses regular gin context (will be wrapped by middleware)
func regularHandler(c *gin.Context) {
	data := map[string]any{
		"id":   "123",
		"name": "John Doe",
	}
	c.JSON(http.StatusOK, data)
}
```
The [response.go](conf/response.go) file contains some standard response methods, such as `Success`, `BadRequest`, `ValidationError` and so on, in case the developer wishes to not state the response status code every time they write a response. They can be used to as follows -

```go
// Handler that uses enhanced context (bypasses middleware wrapping)
func enhancedHandler(c *gin.Context) {
	ec := GetEnhancedContext(c)
	
	data := map[string]any{
		"id":   "456",
		"name": "Jane Doe",
	}
	
	ec.SuccessWithMessage("User retrieved successfully", data)
}
```

It is suggested to use the `ValidationError` function to create the response object in case of validation errors for `POST` and `PATCH` requests.

In case you wish to use both standard and enchanced context responses, make sure you use a `return` statement every time you use an enchanced context response, if it is not the final response of the function, when the function should be exited. Such use case can be as follows -

```go
// Handler that uses both standard and enhanced context
func standardAndenchancedHandler(c *gin.Context) {
    ec := GetEnhancedContext(c)

    var body struct {
        Name string
        Email string
    }

    if err := c.ShouldBindJSON(&body); err != nil {
		ec.ValidationError(err.Error())
		return
	}

    c.JSON(http.StatusOK, body)
}
```

### Extracting Objects From The Request Context
In case of authenticated APIs, the `User` object is stored in the request context. You can retrieve the object by using the `ExtractUserFromContext` function from the `utils` module.

In case of the project, team and bug endpoints, where the IDs are set as path parameters, the `Project` and `Bug` objects are also stored in the request context. You can retrieve them by using the `ExtractProjectFromContext` and `ExtractBugFromContext` functions, respectively, from the `utils` module.