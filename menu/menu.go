package menu

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IrvinTM/frenia/crypt"
	"github.com/IrvinTM/frenia/model"
	"github.com/IrvinTM/frenia/types"
	"github.com/IrvinTM/frenia/util"
)

func Initial() {
	
	if len(os.Args) > 1 {
		args := os.Args[1:]
		switch args[1] {
		case "add":

		case "get":
			if len(os.Args) > 2 {
				var key string
				fmt.Println("Enter your key")
				fmt.Scanln(&key)
				pass, err := model.Read(key,os.Args[2])
				if err != nil {
					fmt.Println(err)
				}else{
					fmt.Printf("Password: %s", pass)
				}
			}


		default:
			fmt.Println("Invalid arg")

		}
	}
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
		passwords := types.PasswordDB{Passwords: map[string]string{"":""}}

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
			fmt.Println("Your passwords:")
			for account := range db.Passwords {
				fmt.Printf("%s\n", account)
			}
		case "2":
			var account, password string
			fmt.Println("Cuenta:")
			_, err := fmt.Scanln(&account)
			if err != nil {
				fmt.Printf("Error scanning line %v", err.Error())
			}
			fmt.Println("Password:")
			_, err = fmt.Scanln(&password)
			if err != nil {
				fmt.Printf("Error scanning line %v", err.Error())
			}
			model.Save(key, account,password)
		case "0":
			return
		default:
			fmt.Println("Invalid option")
		}

	}
}
