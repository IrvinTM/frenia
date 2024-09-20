package menu

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IrvinTM/gopass/crypt"
	"github.com/IrvinTM/gopass/handlers"
	"github.com/IrvinTM/gopass/types"
	"github.com/IrvinTM/gopass/util"
)

func Initial(){
	art := `

________________________________ _______  .___   _____   
\_   _____/\______   \_   _____/ \      \ |   | /  _  \  
 |    __)   |       _/|    __)_  /   |   \|   |/  /_\  \ 
 |     \    |    |   \|        \/    |    \   /    |    \
 \___  /    |____|_  /_______  /\____|__  /___\____|__  /
     \/            \/        \/         \/            \/ 

	 `
	var key string
	DbPath := util.GetDbPath()
	var db types.PasswordDB
	if util.CheckFileExists(DbPath) {
		fmt.Println("Found database")
	} else {
		fmt.Println("DataBase not found creiting one...")
		fmt.Println(DbPath)
		fmt.Println("Please enter your key")
		fmt.Scanln(&key)
		passwords := types.PasswordDB{[]types.Password{{Account: "", Password: ""}}}

		text, err := json.Marshal(passwords)
		if err != nil {
			log.Fatalf("error %v", err.Error())
		}
		crypt.Encrypt(key, string(text), DbPath)

	}
	for {
		fmt.Println(art)
		fmt.Println("Ingrese una opcion")
		fmt.Println("1. Abrir mi Baul")
		fmt.Println("2. Agregar contrasena")
		fmt.Println("0. Salir")

		var option string

		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("ingrese su key")
			fmt.Scanln(&key)

			fmt.Println("tratado de desencriptar")
			decoded := crypt.Decrypt(DbPath, key)
			err := json.Unmarshal(decoded, &db)
			if err != nil {
				log.Fatalf("there was an error err: %v", err.Error())
			}
			for i, pass := range db.Passwords{
				fmt.Println("Your passwords:")
				fmt.Printf("%v %s",i+1, pass.Account)
			}
		case "2":
			var NewPassword types.Password
			fmt.Println("Cuenta:")
			_, err := fmt.Scanln(&NewPassword.Account)
			if err != nil {
				fmt.Printf("Error scanning line %v", err.Error())
			}
			fmt.Println("Password:")
			_, err = fmt.Scanln(&NewPassword.Password)
			if err != nil {
				fmt.Printf("Error scanning line %v", err.Error())
			}
			handlers.AddPassword(key, NewPassword)
		case "0":
			return
		default:
			fmt.Println("Invalid option")
		}

	}
}