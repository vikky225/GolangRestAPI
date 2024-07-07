# REST_API Package

This package contains the REST API implementation for the Golang project.

## Endpoints

- `GET /events`: Get all events
- `GET /events/:id`: Get a specific event
- `POST /events`: Create a new event
- `PUT /events/:id`: Update an event
- `DELETE /events/:id`: Delete an event
- `POST /events/:id/register`: Register for an event
- `DELETE /events/:id/register`: Cancel event registration
- `POST /signup`: User signup
- `POST /login`: User login

## Authentication

- Routes are protected using JWT authentication
- User signup and login functionality included

## Usage

1. Initialize the database by running `db.InitDB()` in the [main](cci:1:///Users/vikasmalviya/REST_API/main.go:10:0-19:1) package.
2. Register the routes using `routes.RegisteredRoutes(server)` in the [main](cci:1:///Users/vikasmalviya/REST_API/main.go:10:0-19:1) package.

## Dependencies

- `github.com/gin-gonic/gin`: Gin web framework
- `github.com/mattn/go-sqlite3`: SQLite database driver
- `github.com/golang-jwt/jwt/v5`: JWT library

## Author

- Vikas Malviya

Feel free to reach out for any questions or contributions related to the REST_API package!
