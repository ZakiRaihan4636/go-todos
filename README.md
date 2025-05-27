# 📝 Go Todo API

A simple RESTful API for a Todo List application built with **Golang** and the **Echo Framework**.

---

## 🚀 Features

- ✅ Create Todo
- ✅ Read Todos
- ✅ Update Todo
- ✅ Delete Todo
- ✅ Mark as Done / Undone

---

## 🛠 Tech Stack

- [Golang](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- [MySQL / MariaDB](https://www.mysql.com/)

---

## ⚙️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/zakiraihan4636/go-todos.git
cd go-todos
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Database

```bash
CREATE DATABASE go_todos;

CREATE TABLE todos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  done TINYINT NOT NULL DEFAULT 0
);
```

### 4. Run the Server

```bash
go run main.go
```

---

## 📝 License

This project is released under the [MIT License](https://opensource.org/licenses/MIT).

Copyright (c) 2023 [Zaki Raihan](https://github.com/zakiraihan4636)

---

[![Star](https://img.shields.io/github/stars/zakiraihan4636/go-todos?style=social)](https://github.com/zakiraihan4636/go-todos)

[![Fork](https://img.shields.io/github/forks/zakiraihan4636/go-todos?style=social)](https://github.com/zakiraihan4636/go-todos/fork)

[![Watch](https://img.shields.io/github/watchers/zakiraihan4636/go-todos?style=social)](https://github.com/zakiraihan4636/go-todos/watchers)
