# 🧩 Go Board — Simple Kanban Task Manager (Go + Vanilla JS)

A lightweight **Kanban-style task board** built with **Go (Golang)** on the backend and **HTML + JavaScript** on the frontend.  
Supports creating, filtering, and organizing tasks visually by status: **To Do**, **In Progress**, and **Done**.

---

##
---

## ⚙️ Requirements

- Go 1.21+
- PostgreSQL (or any SQL database with minor changes)
- Modern browser (Chrome, Edge, Safari)

---

## 🚀 Setup & Run

### 1️⃣ Clone the repo

```bash
git clone https://github.com/yourusername/go-board-app.git
cd go-board-app
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Configure environment

Create a `.env` file in the root directory:

```bash
DATABASE_URL=postgres://user:password@localhost:5432/board_db?sslmode=disable
PORT=8080
```

### 4️⃣ Run the backend server

```bash
go run ./cmd/server
```

Your API will run at:  
👉 **http://localhost:8080**

---

## 🧠 API Endpoints

| Method | Endpoint           | Description                 |
|--------|--------------------|-----------------------------|
| GET    | `/boards`          | Get all boards              |
| GET    | `/boards?user_id=X`| Get boards filtered by user |
| POST   | `/boards`          | Create new board/task       |

### Example: Create Board
```bash
POST /boards
Content-Type: application/json

{
  "name": "Implement Drag-and-Drop",
  "description": "Add frontend drag feature",
  "owner_user_id": "User 1",
  "status": "todo"
}
```

---

## 💻 Frontend

Located in:  
`/web/index.html`

### Features
- Create new tasks
- Filter tasks by user
- Drag and drop between columns
- Real-time UI refresh after create/filter

Run locally by simply opening in your browser:
```
open web/index.html
```

Make sure your backend API is running at `http://localhost:8080`.

---

## 🧩 Example Screenshot

```
📝 To Do       🚧 In Progress       ✅ Done
[Add backend]  [Connect DB]         [Setup CORS]
[Make filter]  [Fix drag-drop]
```

---

## 🛠️ Technologies Used

- **Backend:** Go (Gin or net/http)
- **Database:** PostgreSQL
- **Frontend:** HTML5, CSS3, Vanilla JS (no framework)
- **Architecture:** Clean Architecture + Feature Modules

---

## 🔒 CORS Setup

In `internal/middleware/cors.go`:
```go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}
```

---

## 📦 Future Improvements

- [ ] Persistent drag-and-drop status updates  
- [ ] User authentication  
- [ ] Realtime updates with WebSocket  
- [ ] Dark mode  

---

## 👨‍💻 Author

**Fauzan Habiburrohman**  
Developer Engineer — 7 years experience  
📧 your.email@example.com  
🌐 your-portfolio-or-github-link
