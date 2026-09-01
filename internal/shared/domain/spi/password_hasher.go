package spi

// PasswordHasher defines the contract to hash and verify passwords, shared by any module that manages user credentials.
type PasswordHasher interface {
	// Hash returns a salted hash of the given plain-text password.
	Hash(password string) (string, error)
	// Verify reports whether the given plain-text password matches the given hash.
	Verify(hash string, password string) bool
}
