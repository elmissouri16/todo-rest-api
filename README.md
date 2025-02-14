# Todo REST API

> **Note**: This is my first Go project, created for learning purposes and is currently a work in progress.

A simple REST API built with Go, Gin framework, and SQLite database for managing todo items.

## Technologies Used

- Go
- Gin Web Framework
- GORM (with SQLite driver)

## Setup

1. Clone the repository
2. Make sure you have Go installed
3. Install dependencies:

```bash
go mod init todo-rest-api
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

4. Run the application:

```bash
go run main.go
```

The server will start on port 9090.

## API Endpoints

### Health Check

- `GET /ping`
  - Response: "pong"

### Todo Operations

- `GET /todos`

  - Returns all todo items
  - Response: Array of todo objects

- `GET /todos/add?content=<todo_content>`
  - Creates a new todo item
  - Query Parameter: content
  - Response: Created todo object

## Data Structure

Todo object:

```json
{
  "ID": integer,
  "CreatedAt": timestamp,
  "UpdatedAt": timestamp,
  "DeletedAt": timestamp,
  "Content": string
}
```
