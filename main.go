package main

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"
)

/* User struct */
type User struct {
	ID    int
	Name  string
	Email string
}

var users []User
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== USER MANAGEMENT SYSTEM =====")
		fmt.Println("1. Create User")
		fmt.Println("2. Read User by ID")
		fmt.Println("3. Update User")
		fmt.Println("4. Delete User")
		fmt.Println("5. List All Users")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		choiceInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue
		}

		choiceInput = strings.TrimSpace(choiceInput)
		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		switch choice {
		case 1:
			createUser(reader)
		case 2:
			readUser(reader)
		case 3:
			updateUser(reader)
		case 4:
			deleteUser(reader)
		case 5:
			listUsers()
		case 6:
			fmt.Println("Exiting application. Goodbye 👋")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func isUnique(name, email string) bool {
	for _, user := range users {
		if user.Name == name {
			fmt.Println("❌ Name already exists")
			return false
		}
		if user.Email == email {
			fmt.Println("❌ Email already exists")
			return false
		}
	}
	return true
}

/* CREATE USER */
func createUser(reader *bufio.Reader) {
	fmt.Print("Enter Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name")
		return
	}
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty")
		return
	}

	fmt.Print("Enter Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading email")
		return
	}
	email = strings.TrimSpace(email)
	if !isUnique(name, email) {
		return
	}
	if !isValidEmail(email) {
		fmt.Println("❌ Valid email ID required")
		return
	}

	user := User{
		ID:    nextID,
		Name:  name,
		Email: email,
	}

	users = append(users, user)
	nextID++

	fmt.Println("✅ User created successfully!")
}

/* READ USER BY */
func readUser(reader *bufio.Reader) {
	fmt.Println("\nRead user by:")
	fmt.Println("1. ID")
	fmt.Println("2. Name")
	fmt.Println("3. Email")
	fmt.Print("Choose option (1-3): ")

	optionInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading option")
		return
	}
	option := strings.TrimSpace(optionInput)

	switch option {

	case "1":
		fmt.Print("Enter User ID: ")
		idInput, _ := reader.ReadString('\n')
		idInput = strings.TrimSpace(idInput)

		id, err := strconv.Atoi(idInput)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		for _, user := range users {
			if user.ID == id {
				printUser(user)
				return
			}
		}

	case "2":
		fmt.Print("Enter User Name: ")
		nameInput, _ := reader.ReadString('\n')
		name := strings.ToLower(strings.TrimSpace(nameInput))

		for _, user := range users {
			if strings.ToLower(user.Name) == name {
				printUser(user)
				return
			}
		}

	case "3":
		fmt.Print("Enter User Email: ")
		emailInput, _ := reader.ReadString('\n')
		email := strings.ToLower(strings.TrimSpace(emailInput))

		for _, user := range users {
			if strings.ToLower(user.Email) == email {
				printUser(user)
				return
			}
		}

	default:
		fmt.Println("Invalid option")
		return
	}

	fmt.Println("❌ User not found")
}

func printUser(user User) {
	fmt.Println("\n--- User Details ---")
	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
}

/* UPDATE USER */
func updateUser(reader *bufio.Reader) {
	fmt.Print("Enter User ID to Update: ")
	idInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading ID")
		return
	}

	idInput = strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i := range users {
		if users[i].ID == id {

			fmt.Print("Enter New Name: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading name")
				return
			}
			name = strings.TrimSpace(name)

			fmt.Print("Enter New Email: ")
			email, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading email")
				return
			}
			email = strings.TrimSpace(email)

			if name == "" || email == "" {
				fmt.Println("Name and Email cannot be empty")
				return
			}

			users[i].Name = name
			users[i].Email = email

			fmt.Println("✅ User updated successfully!")
			return
		}
	}

	fmt.Println("❌ User not found")
}

/* DELETE USER */
func deleteUser(reader *bufio.Reader) {
	fmt.Print("Enter User ID to Delete: ")
	idInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading ID")
		return
	}

	idInput = strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Println("✅ User deleted successfully!")
			return
		}
	}

	fmt.Println("❌ User not found")
}

/* LIST USERS */
func listUsers() {
	if len(users) == 0 {
		fmt.Println("No users available")
		return
	}

	fmt.Println("\n--- All Users ---")
	for _, user := range users {
		fmt.Printf("ID: %d | Name: %s | Email: %s\n",
			user.ID, user.Name, user.Email)
	}
}
