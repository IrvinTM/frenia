package main

import (
	"fmt"

	"github.com/IrvinTM/gopass/crypt"
)

type passwordDB struct {
	Passwords []password `json:"passwords"`
}

type password struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func main() {

	var key string
	filename := "database"
	// 	passwords := passwordDB{[]password{{Account: "facebook", Password: "123456"},{Account: "google", Password: "micontrasenadegoogle"}}}

	// text, err := json.Marshal(passwords)
	// if err != nil {
	// 	log.Fatalf("error %v", err.Error())
	// }

	art := `

________________________________ _______  .___   _____   
\_   _____/\______   \_   _____/ \      \ |   | /  _  \  
 |    __)   |       _/|    __)_  /   |   \|   |/  /_\  \ 
 |     \    |    |   \|        \/    |    \   /    |    \
 \___  /    |____|_  /_______  /\____|__  /___\____|__  /
     \/            \/        \/         \/            \/ 

	 `

	for{
		fmt.Println(art)
		fmt.Println("Ingrese una opcion")
		fmt.Println("1. Abrir mi vaul")
		fmt.Println("1. Salir")

		var option string

		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("ingrese su key")
			fmt.Scanln(&key)

			// crypt.Encrypt(key, string(text), filename)

			fmt.Println("tratado de desencriptar")
			crypt.Decrypt(filename, key)
			return
		case "0":
			return
			
		}

	}

}
