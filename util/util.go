package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/IrvinTM/frenia/config"
	"github.com/IrvinTM/frenia/types"
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



// This fuc is used to get the config file and save it to the global config object
// if the file and config folder does not exist, it creates the folders and files in
// home/.config/frenia/config.json
func CheckConfigFile() bool {
	if config.GlobalConfig.Mode == "debug" {
		fmt.Println("checking config file")
	}
	if configPath, err := GetConfigPath(); err == nil {
		fileContent := ReadFile(configPath)
		if err := json.Unmarshal(fileContent, &config.GlobalConfig); err == nil {
			fmt.Println("config file read success")
			fmt.Println(config.GlobalConfig)
		}
		//TODO 2026-07-11T04:37:19.066Z
		// Read mode from args
		// Read/Save salt
		// Test/Unlock db

	return true
	}
		fmt.Println("Config file not found creating one")
		configPath := path.Join(GetHomeDir()+"/.config", "frenia")
		err := os.MkdirAll(configPath, 0755)
		if err != nil{
			fmt.Printf("There was an error creating the config folder: \n %s \n", err)
			panic("Error creating config folder")
		}

		file, err := os.OpenFile(path.Join(configPath, "config.json"), os.O_CREATE|os.O_WRONLY, 0)
		if err != nil{
			fmt.Printf("There was an error creating the config file: \n %s \n", err)
			panic("Error creating config file")
		}
		defer file.Close()
		return false
}

func UpdateConfig(config *types.Configuration){
	if configPath, err := GetConfigPath(); err == nil {
		if jsonConfig, err := json.Marshal(config); err == nil {
			os.WriteFile(configPath, jsonConfig, 0644)
		}
	}
}

