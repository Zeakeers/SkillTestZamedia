package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

var idCounter int

func NewUser(name, email string) User {
	idCounter++
	return User{
		ID:    idCounter,
		Name:  name,
		Email: email,
	}
}

func PrintUserDetails(u *User) {
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
}

func main() {
	fmt.Println("=== Test Struktur User ===")
	user1 := NewUser("Alice", "alice@example.com")
	user2 := NewUser("Bob", "bob@example.com")
	user3 := NewUser("Charlie", "charlie@example.com")

	PrintUserDetails(&user1)
	PrintUserDetails(&user2)
	PrintUserDetails(&user3)
}
