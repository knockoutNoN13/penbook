package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func generatePasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hashedPassword)
}

func main() {
	fmt.Printf(generatePasswordHash("admin"))
}
