package user

import (
	// Import the sql package for interacting with the SQL database.
	"database/sql"
	"fmt"

	// Import the types package, which likely contains data types used across the application, such as User.
	"github.com/FreekAlberti/Ecom/cmd/types"
)

// Store struct represents the data store that interacts with the user data in the database.
// It holds a reference to the SQL database connection.
type Store struct {
	db *sql.DB // SQL database connection.
}

// NewStore is a constructor function that initializes and returns a new instance of Store.
// It accepts a pointer to a sql.DB, which represents the database connection.
func NewStore(db *sql.DB) *Store {
	// Return a new instance of Store with the provided database connection.
	return &Store{db: db}
}

// GetUserByEmail is a method on the Store struct that retrieves a user from the database by their email address.
// It returns a pointer to a User object and an error if something goes wrong.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// Execute an SQL query to find the user by email.
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		// Return nil and the error if the query execution fails.
		return nil, err
	}

	// Initialize a new User object to store the retrieved data.
	u := new(types.User)
	// Iterate over the result set to populate the User object.
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			// Return nil and the error if scanning the row fails.
			return nil, err
		}
	}

	// Check if the User object has a valid ID.
	// If not, it means the user was not found, and an error is returned.
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	// Return the populated User object and nil (no error).
	return u, nil
}

// scanRowIntoUser is a helper function that scans a row from the result set into a User object.
// It returns a pointer to the User object and an error if something goes wrong.
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	// Initialize a new User object to store the scanned data.
	user := new(types.User)

	// Scan the columns from the current row into the User object's fields.
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	// If scanning fails, return nil and the error.
	if err != nil {
		return nil, err
	}

	// Return the populated User object and nil (no error).
	return user, nil
}

// GetUserByID is a method on the Store struct that retrieves a user from the database by their ID.
// The method is not yet implemented and currently returns nil and no error.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	// Implementation is needed to return a user by ID.
	return nil, nil
}

// CreateUser is a method on the Store struct that adds a new user to the database.
// The method is not yet implemented and currently returns no error.
func (s *Store) CreateUser(user types.User) error {
	// Implementation is needed to create a new user in the database.
	return nil
}
