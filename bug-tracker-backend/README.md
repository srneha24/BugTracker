# BUG TRACKER BACKEND

Instructions on how to manage the backend of the project is explained below. The instructions presume that you are already in the _bug-tracker-backend_ folder when you're running commands. So, before doing anything, naviagate to that folder by running `cd bug-tracker-backend` on your terminal.

## Prerequisites
- Go >= 1.24.3
- Postgres >= 16.0

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

Run the following command on your terminal

```bash
go run cmd/main.go
```

To stop the server, press `Ctrl+C`.