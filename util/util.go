package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/IrvinTM/frenia/config"
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

func GetConfigPath() (string, error) {
	configPath := path.Join(GetHomeDir()+"/.config", "frenia", "config.json")
	if CheckFileExists(configPath) {
		fmt.Printf("Config exists at: %s \n", configPath)
		return configPath, nil
	}
	return "", errors.New("Config path was not found")
}

// this fuc will be used to read the config file and check if debug mode and so on
func CheckConfigFile() {
	if path, err := GetConfigPath(); err == nil {
		fileContent := ReadFile(path)
		if err := json.Unmarshal(fileContent, &config.GlobalConfig); err == nil {
			fmt.Println("config file read success")
			fmt.Println(config.GlobalConfig)
		}
		fmt.Println(err)
		//TODO 2026-07-11T04:37:19.066Z
		// Create the file
		// Read mode from args
		// Read/Save salt
		// Test/Unlock db

	}
}
