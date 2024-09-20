package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path"

	"github.com/IrvinTM/gopass/crypt"
	"github.com/IrvinTM/gopass/util"
)

type passwordDB struct {
	Passwords []password `json:"passwords"`
}

type password struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func main() {
	art := `

________________________________ _______  .___   _____   
\_   _____/\______   \_   _____/ \      \ |   | /  _  \  
 |    __)   |       _/|    __)_  /   |   \|   |/  /_\  \ 
 |     \    |    |   \|        \/    |    \   /    |    \
 \___  /    |____|_  /_______  /\____|__  /___\____|__  /
     \/            \/        \/         \/            \/ 

	 `
	var key string
	homeDir := util.GetHomeDir()
	filename := path.Join(homeDir+"/frenia", "database")
	var db passwordDB
	if util.CheckFileExists(filename) {
		fmt.Println("Found database")
	} else {
		fmt.Println("DataBase not found creiting one...")
		fmt.Println(filename)
		fmt.Println("Please enter your key")
		fmt.Scanln(&key)
		passwords := passwordDB{[]password{{Account: "", Password: ""}}}

		text, err := json.Marshal(passwords)
		if err != nil {
			log.Fatalf("error %v", err.Error())
		}
		crypt.Encrypt(key, string(text), filename)

		
		for {
			fmt.Println(art)
			fmt.Println("Ingrese una opcion")
			fmt.Println("1. Abrir mi Baul")
			fmt.Println("0. Salir")

			var option string

			fmt.Scanln(&option)

			switch option {
			case "1":
				fmt.Println("ingrese su key")
				fmt.Scanln(&key)

				fmt.Println("tratado de desencriptar")
				decoded := crypt.Decrypt(filename, key)
				err := json.Unmarshal(decoded, &db)
				if err != nil {
					log.Fatalf("there was an error err: %v", err.Error())
				}
				fmt.Printf("tus contrasenas son para %s es %v\n", db.Passwords[0].Account, db.Passwords[0].Password)
			case "0":
				return

			}

		}

	}
}
