# go-graphql-sqlc-api
# GraphQL + SQLC + GoLang: Users and Social Accounts

This is a fully working GraphQL API built using **GoLang**, **gqlgen**, **PostgreSQL**, and **SQLC**. It showcases how to efficiently fetch users and their associated social media accounts while avoiding N+1 problems using a single optimized SQL join.

---

## 📦 Tech Stack

- 🧬 GraphQL (via `gqlgen`)
- 💡 GoLang
- 🐘 PostgreSQL
- ⚙️ SQLC for type-safe SQL generation

---

---

## 🔧 Getting Started (Without Docker)

### 1. Set up PostgreSQL Locally

Ensure PostgreSQL is installed and running.

```bash
createdb demo
```

> Alternatively, connect using `psql` or a GUI (like pgAdmin) and create a database named `demo`.

### 2. Load the SQL Schema

```bash
psql -U postgres -h localhost -d demo -f sql/schema.sql
```

### 3. Setup Environment Variables

Edit `.env` file as needed:

```env
DB_USER=user
DB_PASSWORD=user_password
DB_NAME=db
DB_HOST=localhost
DB_PORT=5432
```

### 4. Install Dependencies

```bash
go get github.com/joho/godotenv
```

### 5. Generate Code

```bash
sqlc generate
go run github.com/99designs/gqlgen generate
```

### 6. Run the Application

```bash
go run main.go
```

Then visit:  
➡️ `http://localhost:8080/playground`

---

## 🚀 Example Query

```graphql
query {
  users {
    id
    name
    accounts {
      id
      platform
      username
    }
  }
}
```

---

## 🧠 What You Learn

- ✅ How to use SQLC to define typed SQL queries in Go
- ✅ How to structure GraphQL resolvers to return nested data from a flat SQL result
- ✅ How to avoid the N+1 query problem using LEFT JOINs
- ✅ Clean separation of SQL, Go models, and GraphQL layers

---

## 🛠 Tools Required

- Go 1.20+
- PostgreSQL
- [sqlc](https://docs.sqlc.dev/)
- [gqlgen](https://gqlgen.com/)

---

## 🙋‍♂️ Author

Built by Sai Teja Bandamidi – feel free to fork and extend!