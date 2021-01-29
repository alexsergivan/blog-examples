package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	Password string `json:"password" form:"password"`
	Name string `json:"name" form:"name"`
}

// LoadTestUser loads a dummy user.
func LoadTestUser() *User {
	// Just for demonstration purpose, we create a user with the encrypted "test" password.
	// In real-world applications, you might load the user from the database by specific parameters (email, username, etc.)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &User{Password: string(hashedPassword), Name: "Test user"}
}