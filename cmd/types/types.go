package types

import "time"

// UserStore is an interface that defines the contract for any data store that handles user-related operations.
// Any struct that implements these methods can be used as a UserStore in the application.
type UserStore interface {
	// GetUserByEmail retrieves a user by their email address.
	// It returns a pointer to a User object and an error if something goes wrong.
	GetUserByEmail(email string) (*User, error)

	// GetUserByID retrieves a user by their unique ID.
	// It returns a pointer to a User object and an error if something goes wrong.
	GetUserByID(id int) (*User, error)

	// CreateUser adds a new user to the data store.
	// It accepts a User object and returns an error if the operation fails.
	CreateUser(User) error
}

// User struct represents a user in the application.
// It contains various fields such as ID, first name, last name, email, password, and the time the user was created.
type User struct {
	ID        int       `json:"id"`        // Unique identifier for the user.
	FirstName string    `json:"firstName"` // User's first name.
	LastName  string    `json:"lastName"`  // User's last name.
	Email     string    `json:"email"`     // User's email address.
	Password  string    `json:"-"`         // User's hashed password. This field is omitted from JSON responses.
	CreatedAt time.Time `json:"createdAt"` // Timestamp when the user was created.
}

// RegisterUserPayload struct is used to capture and validate the data sent when a new user is registering.
// The struct tags specify JSON keys and validation rules for each field.
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`              // First name is required.
	LastName  string `json:"lastName" validate:"required"`               // Last name is required.
	Email     string `json:"email" validate:"required,email"`            // Email is required and must be a valid email format.
	Password  string `json:"password" validate:"required,min=3,max=130"` // Password is required, with a minimum of 3 and a maximum of 130 characters.
}
