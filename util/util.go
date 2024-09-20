package util

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func GetHomeDir() string{
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