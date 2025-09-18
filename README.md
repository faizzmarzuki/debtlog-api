# debtlog-api

DebtLog is a REST API built with **Go**.  
It’s my personal learning project for backend development, while also serving as a tool to track money debts between me and my friends.

Right now it focuses only on **money debt**, but I plan to expand it later for other types of debts (items, favors, etc).

---

## 🚀 Learning Purpose

This project is a hands-on way for me to explore:

- Building a REST API with **Go** and **Gin**
- Working with relational databases using **GORM** and **PostgreSQL**
- Authentication & Authorization with **JWT**
- Configuration management with `.env`
- Database migrations and ORM models
- Testing in Go

---

## 📌 Features (Current)

- User registration and authentication (JWT-based)
- Create new debts (who owes who, amount, notes)
- Fetch debts (all debts, or specific debt by ID)
- Update or settle a debt
- Delete a debt

---

## 🛠️ Tech Stack

- **Language:** Go (v1.25.1)
- **Framework:** Gin (`github.com/gin-gonic/gin`)
- **ORM:** GORM (`gorm.io/gorm`, `gorm.io/driver/postgres`)
- **Database:** PostgreSQL
- **Auth:** JWT (`github.com/golang-jwt/jwt/v5`)
- **Config:** godotenv (`github.com/joho/godotenv`)

---

## 📂 Project Structure (example)

```
debtlog-api/
├── config/ # Configuration (database, env setup, etc.)
│ └── db.go
│
├── controllers/ # Controllers (handle HTTP requests)
│ ├── auth_controller.go
│ ├── debter_controller.go
│ ├── debtlog_controller.go
│ ├── receipt_controller.go
│ └── share_controller.go
│
├── middleware/ # Custom middlewares
│ └── auth_middlewa re.go
│
├── models/ # Database models (GORM)
│ ├── debt_link.go
│ ├── debt_log.go
│ ├── debt_log_debter.go
│ ├── debter.go
│ ├── receipt.go
│ └── user.go
│
├── routes/ # Route definitions
│ └── routes.go
│
├── tests/ # Unit & integration tests
│ └── auth_test.go
│
├── utils/ # Utility functions
│ └── jwt.go
│
├── .env # Environment variables
├── go.mod
├── go.sum
├── main.go # App entry point
└── main.go.save # Backup file
```

---

## ⚡ Getting Started

### Prerequisites

Before you begin, make sure you have:

- **Go 1.25.1+** installed → [Download Go](https://go.dev/dl/)
- **PostgreSQL** installed and running locally (or on Docker)

---

### 1. Fork or Clone the Repository

You can either fork the project to your own GitHub account, or clone it directly:

```bash
# Clone directly
git clone https://github.com/faizzmarzuki/debtlog-api.git

# or fork first, then clone your fork:
git clone https://github.com/<your-username>/debtlog-api.git

cd debtlog-api
```

### 2. Install Dependencies

Go modules will pull everything you need:

```bash
go mod tidy
```

### 3. Setup Environment Variables

Copy the example environment file:

```bash
cp .env.example .env
```

Then edit `.env` with your PostgreSQL details:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=debtlog
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

### 4. Create the Database

Log into PostgreSQL and create the database:

```bash
psql -U postgres
CREATE DATABASE debtlog;
\q
```

### 5. Run Migrations

If migrations are set up with GORM auto-migrate (inside `main.go`), the tables will be created automatically when you first run the app.

If you use a migration tool like `golang-migrate`, run something like:

```bash
migrate -path migrations -database "postgres://postgres:yourpassword@localhost:5432/debtlog?sslmode=disable" up
```

_(Optional: depends on how you've set it up — by default, this repo auto-migrates via Go code.)_

### 6. Start the Server

```bash
go run main.go
```

Server should now be running at `http://localhost:8080`.

### 7. Run Tests

```bash
go test ./tests -v
```

---

## 📖 API Endpoints

### Public Routes

| Method | Endpoint        | Description                  |
| ------ | --------------- | ---------------------------- |
| POST   | `/register`     | Register a new user          |
| POST   | `/login`        | Login with username/password |
| GET    | `/health`       | Health check                 |
| GET    | `/share/:token` | Access a shareable debt link |

### Protected Routes (JWT required)

#### Debters

| Method | Endpoint       | Description         |
| ------ | -------------- | ------------------- |
| POST   | `/debters`     | Create a new debter |
| GET    | `/debters`     | List all debters    |
| PUT    | `/debters/:id` | Update debter by ID |
| DELETE | `/debters/:id` | Delete debter by ID |

#### Debt Logs

| Method | Endpoint        | Description          |
| ------ | --------------- | -------------------- |
| POST   | `/debtlogs`     | Create a debt log    |
| GET    | `/debtlogs/:id` | Get debt log details |

#### Receipts

| Method | Endpoint                 | Description               |
| ------ | ------------------------ | ------------------------- |
| POST   | `/debtlogs/:id/receipts` | Upload a receipt for debt |

---

## 🛣️ Roadmap

- [ ] Debt categories (money, items, favors)
- [ ] Notifications / reminders
- [ ] User-to-user friend system
- [ ] Swagger documentation
- [ ] Frontend (React/Vue)
- [ ] Mobile app version

---

## ✨ Why This Project?

I wanted a simple but real-world project to practice backend development with Go.
DebtLog is both:

- A **personal tool** to track debts with friends
- A **learning project** to explore Go, Gin, JWT, GORM, and PostgreSQL

---
