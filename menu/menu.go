package menu

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IrvinTM/frenia/config"
	"github.com/IrvinTM/frenia/crypt"
	"github.com/IrvinTM/frenia/model"
	"github.com/IrvinTM/frenia/types"
	"github.com/IrvinTM/frenia/util"
	"golang.org/x/term"
)

func Initial() {

	//Load the config
	configExists := util.CheckConfigFile()
	var pass *string
	
	if configExists {
		fd := int(os.Stdin.Fd())
		fmt.Println("Please enter your password")

		bytepass, err := term.ReadPassword(fd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n Error reading password:", err)
		}

		salt, err := base64.StdEncoding.DecodeString((config.GlobalConfig.Salt))
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n Error decoding salt:", err)
		}
		hashed, err := crypt.HashPassword(string(bytepass), salt)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n Error hashing password:", err)
		}
	// strPass := base64.StdEncoding.EncodeToString([]byte(hashed))

	strPass := hashed
	pass = &strPass

	}else{
		salt := crypt.GenerateRamdomSalt(32)
		configuration := types.Configuration{Salt: base64.StdEncoding.EncodeToString(salt)}
		util.UpdateConfig(&configuration)
	}

	if len(os.Args) > 1 {
		fmt.Print("we got more than 1 arg")
		args := os.Args[1:]
		for i := range len(args){
			fmt.Printf("\n arg %d : %s\n",i, args[i])
		}
		switch args[0] {
		case "add":
			if len(os.Args) > 3 {
				model.Save(*pass, args[1], args[2])
			}
		case "get":
			if len(os.Args) > 2 {
				fmt.Printf("getting pass for %s\n", args[1])
				readPass, err := model.Read(*pass, os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(readPass)
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
	DbPath := util.GetDbPath()
	var db types.PasswordDB
	if util.CheckFileExists(DbPath) {
		fmt.Println("Found database")
	} else {
		fmt.Println("DataBase not found creiting one...")
		fmt.Println(DbPath)
		passwords := types.PasswordDB{Passwords: map[string]string{"": ""}}

		text, err := json.Marshal(passwords)
		if err != nil {
			log.Fatalf("error %v", err.Error())
		}

		fmt.Printf("\n The key being uses is: %d  %s \n",len(*pass), *pass)

		crypt.Encrypt(*pass, string(text), DbPath)

	}
	if len(os.Args) == 1 {
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
			fmt.Println("tratado de desencriptar")
			decoded := crypt.Decrypt(DbPath, *pass)
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
			model.Save(*pass, account, password)
		case "0":
			return
		default:
			fmt.Println("Invalid option")
		}

	}
		
	}
	
}
