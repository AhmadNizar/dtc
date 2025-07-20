# Gajian – Payroll Backend API

This is a Golang REST API service for managing payroll, overtime, attendance, and more.  
It uses the following technologies:

- [Gin](https://github.com/gin-gonic/gin) for HTTP routing  
- [GORM](https://gorm.io/) as ORM for PostgreSQL  
- [Dev Containers](https://containers.dev/) for development setup  
- PostgreSQL as the main database

---

## 📁 Project Structure

```

gajian/
│
├── cmd/
│   └── gajian/
│       └── main.go       # Entry point of the application
|       └── http/
|           └──contollers # API handler for each function
|           └──router     # API route for each function
│
├── internal/             # Internal application layers
│   ├── application/      # Contain logic for each usecase
│   ├── domain/           # Contain interface to source of data (exp: Database)
│   └── dto/              # Contain blueprint of data transfer object
│
├── scripts/              # SQL migration files
├── .devcontainer/        # Dev container configuration
├── go.mod / go.sum
└── README.md

````

---

## 🚀 Getting Started

### 🧱 Prerequisites

- [Docker](https://www.docker.com/)
- [VS Code](https://code.visualstudio.com/) + Dev Containers extension (optional but recommended)

---

### 🐳 Using Dev Container (Recommended)

1. Open the project in VS Code  
2. Run: `Dev Containers: Reopen in Container`  
3. The Go app and PostgreSQL database will be set up automatically.

---

### 💻 Running the App Manually (Outside Dev Container)

```bash
# 1. Copy env vars if necessary
cp .env.example .env

# 2. Start PostgreSQL using Docker
docker-compose up -d

# 3. Run the Go app
go run cmd/gajian/main.go
````

---

### ⚙️ Configuration

Environment variables are handled via `.env` file (or can be passed directly). Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gajian_db
DB_TIMEZONE=Asia/Jakarta
```

---

## 🗄️ Database Migrations

Use [golang-migrate](https://github.com/golang-migrate/migrate) or similar tools to apply SQL migrations in the `/scripts/postgresql/migrations/` folder.

```bash
bash ./scripts/init_psql.sh
```

```bash
bash ./scripts/create_table.sh
```

---

## 🧪 API Testing

Use tools like Postman, cURL, or HTTPie. For example:

```bash
curl http://localhost:8080/api/v1/users
```

---

## 🧰 Dev Tools

* Dev Container runs Go 1.21+ and PostgreSQL 15+
* GORM auto-migrates tables for development
* Migrations are in raw SQL for production compatibility

---

## 📜 License

MIT License
