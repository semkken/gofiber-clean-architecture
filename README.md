# GoFiber Clean Architecture Ticket System by kasemsan k.

A sample ticket system built with Go, Fiber, and Clean Architecture principles.

## Project Structure

```
├── api/
│   ├── dto/          # Data Transfer Objects
│   ├── handlers/     # HTTP Handlers
│   ├── middleware/   # Middleware (e.g., JWT)
│   ├── routes/       # Route definitions
│   └── presenter/    # Response formatting
├── pkg/
│   ├── auth/         # JWT authentication logic
│   ├── commands/     # Command handlers (write use cases)
│   ├── queries/      # Query handlers (read use cases)
│   ├── entities/     # Domain models
│   └── repository/   # Data access abstractions
├── internal/
│   ├── database/     # Database connection
│   └── storage/      # File storage
├── main.go           # Application entry point
└── go.mod            # Go module dependencies
```

## Getting Started

1. **Install Go:**  
   Download and install Go from [https://go.dev/dl/](https://go.dev/dl/).

2. **Install dependencies:**  
   ```
   go mod tidy
   ```

3. **Run the application:**  
   ```
   go run ./main.go
   ```

4. **API Endpoints:**  
   - `POST /tickets` - Create a ticket (JWT required)
   - `POST /tickets/:id/assign` - Assign a ticket (JWT required)
   - `GET /tickets/:id` - Get ticket by ID
   - `GET /tickets` - List all tickets

## Notes

- Update JWT secret key in `main.go` for production use.
- Add your database/storage implementation in `internal/`.
- See each package for more details.

---

Feel free to contribute or

Register
curl -X POST http://localhost:3000/register -H "Content-Type: application/json" -d "{\"username\":\"john_doe\",\"email\":\"john@example.com\",\"password\":\"password123\"}"

Login
curl -X POST http://localhost:3000/login -H "Content-Type: application/json" -d "{\"username\":\"john_doe\",\"password\":\"password123\"}"

Profile
curl -X GET http://localhost:3000/api/profile -H "Authorization: Bearer <token>"

Create Ticket
curl -X POST http://localhost:3000/api/tickets -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d "{\"title\":\"Issue\",\"description\":\"Description of issue\"}"

Assign Ticket
curl -X POST http://localhost:3000/api/tickets/<ticket_id>/assign -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d "{\"assignee_id\":\"user_id\"}"

Get All Tickets
curl -X GET http://localhost:3000/api/tickets -H "Authorization: Bearer <token>"



