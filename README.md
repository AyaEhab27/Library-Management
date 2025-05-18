# 📚 Library Management System in Go

## 🌟 Project Overview
A robust, budget-friendly library management solution developed in Go to help small libraries transition to digital services post-COVID-19. Features dual-server architecture with automatic failover and JSON-based data storage.

## ✨ Key Features

### 📖 Book Operations
| Feature        | Icon | Description                          |
|----------------|------|--------------------------------------|
| Add Books      | ➕   | Add new books to the collection      |
| Search Books   | 🔍   | Find books by ID or title            |
| Sort Books     | 🔄   | Organize by title or publication date|
| Data Persistence| 💾  | Auto-save to JSON files              |

### 👥 Reader Management
| Feature        | Icon | Description                          |
|----------------|------|--------------------------------------|
| Add Readers    | 👤➕ | Register new library members         |
| Remove Readers | ❌   | Delete reader records                |
| Search Readers | 🔎   | Find readers by ID or name           |

### ⚙️ System Architecture
```sh
📦library
├── 📂client       # CLI interface
├── 📂server       # Dual-server implementation
├── 📂models       # Data structures
└── 📂data         # JSON storage


## 🔍 Core Features

### 📖 Book Management
- **Data Fields**: ID, Title, Author, Publication Date, Genre, Publisher, Language  
- **Operations**:
  - ➕ Add new books
  - 🔍 Search by ID/Title
  - 📥📤 Load/Save book data
  - 🔄 Sort by Title/Publication Date

### 👥 Reader Management  
- **Data Fields**: ID, Name, Gender, Birthday, Height, Weight, Employment  
- **Operations**:
  - ➕ Add readers
  - ❌ Remove readers
  - 🔍 Search by ID/Name
  - 👀 View all readers

### ⚡ Server Infrastructure
- 🟢 **Main Server** (Port 8080)  
- 🟠 **Backup Server** (Port 8081) - Auto-activates if main fails  
- 🔄 Health checks every 5 seconds  
- 💾 JSON file storage (No database costs)

### 💻 CLI Interface
- 🎨 Color-coded menus  
- 🔄 Automatic server detection  
- ✅ Intuitive workflows

---

## 🚀 Getting Started

```bash
# Start servers (Main + Backup)
go run main.go server

# Launch CLI application
go run main.go

# Access web interfaces:
# Main:    http://localhost:8080
# Backup:  http://localhost:8081

## 🛠 Technical Specs
- **Language**: Go 1.16+  
- **Storage**: JSON files (No database required)  
- **Ports**: 8080 (Main), 8081 (Backup)  

---

## 🌟 Future Roadmap
- 🌐 Web patron portal  
- 🔖 Book reservation system  
- ⏳ Overdue fine calculations  
- 🔎 Advanced search filters  

---

## 💡 Why This Solution?
- 💰 **Cost-effective** - No database expenses  
- 🛡 **Reliable** - Automatic failover protection  
- 📱 **Modern** - Full digital transformation  
- 👩‍💻 **User-friendly** - Simple CLI and web interfaces  
