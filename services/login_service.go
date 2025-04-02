package services

import (
	"errors"
	"fmt"
	"log"
	"nexttalenta-backend/database"
	"nexttalenta-backend/models"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	userRepo database.UserRepository
}

func NewLoginService(userRepo database.UserRepository) LoginService {
	return LoginService{userRepo: userRepo}
}

func (s LoginService) Authenticate(username, password string) (*models.AuthResponse, error) {
	log.Printf("Authenticating user: %s", username)
	log.Printf("Input Password: [MASKED] %s", password)

	user, err := s.userRepo.FindByUsername(username)
	log.Printf("DB Password: [MASKED] %s", user.Password)
	if err != nil {
		log.Printf("Database error: %v", err)
		return nil, errors.New("invalid username or password")
	}

	log.Printf("Attempt login from user: %s", username)

	if password == user.Password {
		log.Printf("SAMEEE")
	}

	// Check if the password is the default "Default"
	if password == "Default" {
		log.Printf("Login Failed: Password is the default 'Default'")
		return nil, errors.New("your password has been reset to the standard password, please change it now!")
	}

	// Verify password using bcrypt
	if !checkPasswordHash(password, user.Password) {
		log.Printf("Invalid password attempt for user: %s", username)

		remainingAttempts := 3 - (user.FailedAttempts + 1)
		return nil, fmt.Errorf("login failed, wrong password, you have %d remaining attempts", remainingAttempts)
	}

	// Successful authentication
	response := &models.AuthResponse{
		Username:  user.Username,
		IsSuccess: true,
		Message:   "Login successful",
	}

	return response, nil
}

func checkPasswordHash(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
