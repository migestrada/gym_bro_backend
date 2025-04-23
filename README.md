# Gym Bro Backend

This is the backend for the **Gym Bro** application, built with Go and Gin. It provides APIs for managing training plans.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/gym_bro_backend.git
   cd gym_bro_backend

2. Install dependencies::
   ```bash
   go mod tidy

3. Set up the database:
    - Ensure PostgreSQL is running.
    - Update the database credentials in connection/database.go.

4. Run the application::
   ```bash
   go run main.go

---
## Testing
Manual Tests
  -  Use the .http files in tests/manual/ to test the API endpoints with tools like REST Client.

Automated Tests
  ```bash
  go test ./tests/...