# GoFiber Article API

A simple backend service built with **Golang 1.24.3**, **GoFiber**, **GORM**, **Wire**, and **Go-Migrate**, providing RESTful endpoints for managing articles.

---

## Tech Stack
- **Golang v1.24.3**
- **Fiber v2** - HTTP framework
- **GORM** - ORM for MySQL/PostgreSQL
- **Go-Migrate** - Database migration management
- **Google Wire** - Dependency injection
- **PostgreSQL** - Database support

---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/farisfadhail/sv-be.git
cd sv-be
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Configure Database
Set your database connection in `.env` file:
```
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=user
DB_PASS=password
DB_NAME=sv-be
```

### 4. Run Migrations
```bash
make migrate-up
# or for a fresh start
make migrate-fresh
```

### 5. Run the App
```bash
make run
```
App will run at: **http://127.0.0.1:8080**

---

## Makefile Commands

| Command | Description |
|----------|-------------|
| `make run` | Run the main server |
| `make wire` | Generate dependency injection using Google Wire |
| `make migrate-up` | Apply database migrations |
| `make migrate-fresh` | Drop all tables and reapply migrations |
| `make migrate-create name=<migration_name>` | Create a new migration file |

---

## API Endpoints

### Welcome Message
```http
GET http://127.0.0.1:8080
```

### Get All Articles
```http
GET http://127.0.0.1:8080/api/article?limit=5&offset=0
```

### Create a New Article
```http
POST http://127.0.0.1:8080/api/article
Content-Type: application/json

{
  "title": "Pengenalan Web 3.0 dan Blockchain PUBLISH",
  "content": "Web 3.0 merupakan evolusi dari internet yang menekankan pada desentralisasi...",
  "category": "Teknologi",
  "status": "publish"
}
```

### Get Article by ID
```http
GET http://127.0.0.1:8080/api/article/{id}
```

### Update an Article
```http
PUT http://127.0.0.1:8080/api/article/{id}
Content-Type: application/json

{
  "title": "Pengenalan Web 3.0 dan Blockchain (Update)",
  "content": "Dalam pembaruan ini, artikel menjelaskan lebih detail mengenai integrasi...",
  "category": "Teknologi",
  "status": "thrash"
}
```

### Delete an Article
```http
DELETE http://127.0.0.1:8080/api/article/{id}
```

---

## Example Responses

**GET /api/article**
```json
[
  {
    "id": 1,
    "title": "Pengenalan Web 3.0 dan Blockchain",
    "content": "Web 3.0 merupakan evolusi dari internet yang menekankan pada desentralisasi...",
    "category": "Teknologi",
    "status": "publish"
  }
]
```

---

## Project Structure

```
.
├── cmd/
│   ├── seeder.go
│   └── migration.go
├── config/
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── validation/
│   ├── providers/
│   ├── middleware/
│   └── injector/
├── database/
│   ├── seeders/
│   └── migrations/
├── models/
├── resources/
├── utils/
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```
