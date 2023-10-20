package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var password string

	// Prompt the user to enter a password
	fmt.Print("Enter a password: ")
	_, err := fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Generate a salt and hash for the entered password
	hashedPassword, err := hashAndSaltPassword(password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original Password:", password)
	fmt.Println("Hashed Password:", hashedPassword)
}

func hashAndSaltPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Generate the hashed password
	hashedPassword := string(salt)

	return hashedPassword, nil
}
