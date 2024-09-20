package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Encrypt(key string, plainText string, ouputFileName string) {
	dir := filepath.Dir(ouputFileName)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("Directory %s exist skipping", dir)
	}
	strKey := key
	byteSlice := []byte(strKey)
	block := NewAesCipher(byteSlice)
	gcm := NewGcmStream(block)
	encryptedText := GenerateRandomNonce(gcm, []byte(plainText))
	SaveCipherText(encryptedText, ouputFileName)
}

func ReadFile(filename string) []byte {
	plainText, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file err: %v", err.Error())
	}
	return plainText
}

func NewAesCipher(key []byte) cipher.Block {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("There was an error creating the cipher err: %v", err.Error())
	}
	return block
}

func NewGcmStream(block cipher.Block) cipher.AEAD {
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("There was an error in the GCM err: %v", err.Error())
	}
	return gcm
}

func GenerateRandomNonce(gcm cipher.AEAD, plainText []byte) []byte {
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce error err: %v", err.Error())
	}
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return cipherText
}

func SaveCipherText(cipherText []byte, fileName string) {
	err := os.WriteFile(fileName, cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
	fmt.Printf("saved to %s", fileName)
}
