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
