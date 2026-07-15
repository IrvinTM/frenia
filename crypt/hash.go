package crypt

import (
	"crypto/rand"
	"encoding/base64"

	//"fmt"
	"golang.org/x/crypto/argon2"
)

const (
	// Parameters for Argon2id
	timeCost    = 1         // Number of iterations
	memoryCost  = 64 * 1024 // 64MB
	parallelism = 4         // Number of threads
	keyLength   = 32        // Desired key length (32 bytes for AES-256)
)

// HashPassword generates a 32-byte key from the given password using Argon2id.
// It returns the derived key (hash) and the generated salt.
func HashPassword(password string, salt []byte) (string, error) {

	// Generate the Argon2id hash (derived key)
	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLength)

	// Encode the hash and salt in base64 for storage
	encodedHash := base64.StdEncoding.EncodeToString(hash)

	return encodedHash, nil
}

// I need to save the salt

//This will be done in a configuration file

// hash the entered password

func GenerateRamdomSalt(saltLength int) []byte{
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt;
}
