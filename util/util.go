package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
)

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func ReadFile(filename string) []byte {
	plainText, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file err: %v", err.Error())
	}
	return plainText
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return homeDir
}

func GetDbPath() string {
	dbPath := path.Join(GetHomeDir()+"/frenia", "database")
	return dbPath
}
