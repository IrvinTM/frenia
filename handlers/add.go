package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/IrvinTM/gopass/crypt"
	"github.com/IrvinTM/gopass/types"
	"github.com/IrvinTM/gopass/util"
)

func AddPassword(key string, password types.Password){
	var db types.PasswordDB

	data := crypt.Decrypt(util.GetDbPath(), key)
	
	err := json.Unmarshal(data, &db)
	if err != nil {
		fmt.Printf("Error when getting passwords %v", err.Error())
	}
	db.Passwords = append(db.Passwords, password)

	unmarshalled, err := json.Marshal(db)
	if err != nil {
		fmt.Printf("There was an error saving %v", err.Error())
	}
	crypt.Encrypt(key, string(unmarshalled), util.GetDbPath())
}