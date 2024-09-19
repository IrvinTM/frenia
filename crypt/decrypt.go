package crypt

import (
	"crypto/cipher"
	"fmt"
	"log"
)


func Decrypt(filename string, key string){
	de := NewAesCipher([]byte(key))
	AEAD := NewGcmStream(de)
	encriptedBytes := ReadFile(filename)
	decoded, _ := Decipher(encriptedBytes, filename, AEAD )
	fmt.Printf("tus conrasenas son: %v", string(decoded))
}

func Decipher(fileContent []byte, fileName string, gcm cipher.AEAD) ([]byte,error){
	nonce := fileContent[:gcm.NonceSize()]
	fileContent =fileContent[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, fileContent, nil)
	if err != nil {
		log.Fatalf("error decrypting err: %v", err.Error())
		return nil, err
	}
	return plainText, nil
}