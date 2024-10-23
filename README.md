# Run the app
```bash
go mod tidy
go run main.go
```
# Environment configuration example
create config.env file :

```bash
DB_HOST=localhost
DB_NAME=database
DB_PORT=8080
DB_USER=user
DB_PASS=admin

GRPC_PORT=8000

GIN_HOST=localhost
GIN_PORT=800
```