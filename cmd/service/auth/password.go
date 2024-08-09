package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword is a utility function that takes a plain text password as input,
// hashes it using the bcrypt algorithm, and returns the hashed password as a string.
// If an error occurs during the hashing process, it returns an empty string and the error.
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword hashes the password using the bcrypt algorithm.
	// The password is converted to a byte slice, and bcrypt.DefaultCost is used to determine the hashing complexity.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// If there is an error during the hashing process, return an empty string and the error.
		return "", err
	}

	// Convert the hashed password (byte slice) back to a string and return it.
	return string(hash), nil
}
