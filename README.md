# golang-crud-cli
A simple command-line User Management System built in Go (Golang) that supports creating, reading, updating, deleting, and listing users with input validation and email verification.
# Go User Management CLI

A simple **Command-Line User Management System** built using **Golang**.  
This application demonstrates CRUD operations with input validation, unique checks, and email verification.

## 🚀 Features

- Create users with name and email
- Read users by:
  - ID
  - Name
  - Email
- Update user details
- Delete users
- List all users
- Email format validation
- Prevent duplicate names and emails
- Interactive CLI menu

## 🛠️ Tech Stack

- Go (Golang)
- Standard Library only

## 📂 Project Structure
├── main.go
└── README.md

## ▶️ How to Run

1. Install Go (1.18+ recommended)
2. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-user-management-cli.git
3.Navigate into the project:
cd go-user-management-cli
4 . Run the program:
go run main.go
📸 Sample Menu
===== USER MANAGEMENT SYSTEM =====
1. Create User
2. Read User by ID
3. Update User
4. Delete User
5. List All Users
6. Exit
📌 Notes

Data is stored in memory (no database)

Ideal for beginners learning Go CLI applications and CRUD logic
