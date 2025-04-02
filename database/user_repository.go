package database

import (
	"database/sql"
	"errors"
	"nexttalenta-backend/models"
	"time"
)

// UserRepository interface
type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
}

// userRepository struct
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// FindByUsername retrieves a user by username
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, email, name, failed_attempt, password_expiry FROM userdata WHERE username = ?"
	row := r.db.QueryRow(query, username)

	// Declare a temporary string for password_expiry
	var passwordExpiryStr string
	user := &models.User{}

	// Scan the results into the struct fields
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		new(int),           // failed_attempt (ignored for now)
		&passwordExpiryStr, // Scan as string instead of time.Time
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Convert passwordExpiryStr (string) to time.Time
	if passwordExpiryStr != "" {
		user.PasswordExpiry, err = time.Parse("2006-01-02", passwordExpiryStr)
		if err != nil {
			user.PasswordExpiry = time.Time{} // Default to zero time if parsing fails
		}
	}

	return user, nil
}
