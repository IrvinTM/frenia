
package crypt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const (
	// Parameters for Argon2id
	timeCost    = 1   // Number of iterations
	memoryCost  = 64 * 1024  // 64MB
	parallelism = 4   // Number of threads
	saltLength  = 16  // Salt length in bytes
	keyLength   = 32  // Desired key length (32 bytes for AES-256)
)

// HashPassword generates a 32-byte key from the given password using Argon2id.
// It returns the derived key (hash) and the generated salt.
func HashPassword(password string) (string, string, error) {
	// Generate a random salt
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", "", err
	}

	// Generate the Argon2id hash (derived key)
	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLength)

	// Encode the hash and salt in base64 for storage
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)

	return encodedHash, encodedSalt, nil
}


