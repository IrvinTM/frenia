package model

import (
	"encoding/json"
	"fmt"

	"github.com/IrvinTM/frenia/crypt"
	"github.com/IrvinTM/frenia/types"
	"github.com/IrvinTM/frenia/util"
)

func Save(key string, account string, password string ) {
	var db types.PasswordDB

	data := crypt.Decrypt(util.GetDbPath(), key)

	err := json.Unmarshal(data, &db)
	if err != nil {
		fmt.Printf("Error when getting passwords %v", err.Error())
	}
	db.Passwords[account] = password

	unmarshalled, err := json.Marshal(db)
	if err != nil {
		fmt.Printf("There was an error saving %v", err.Error())
	}
	crypt.Encrypt(key, string(unmarshalled), util.GetDbPath())
}

func Delete(key string, account string){
	var db types.PasswordDB

	data := crypt.Decrypt(util.GetDbPath(), key)

	err := json.Unmarshal(data, &db)
	if err != nil {
		fmt.Printf("Error when getting passwords %v", err.Error())
	}

	delete(db.Passwords, account)

	unmarshalled, err := json.Marshal(&db)
	if err != nil {
		fmt.Printf("There was an error saving %v", err.Error())
	}
	crypt.Encrypt(key, string(unmarshalled), util.GetDbPath())
}

func Read(key string, account string) (string, error){
	var db types.PasswordDB
	var password string
	var e error

	data := crypt.Decrypt(util.GetDbPath(), key)

	err := json.Unmarshal(data, &db)
	if err != nil {
		fmt.Printf("Error when getting passwords %v", err.Error())
	}
	element, exist := db.Passwords[account]
	if exist{
		password = element
	}else{
		e = fmt.Errorf("the account %s does not exist in the database", account)
	}
	return password, e
}