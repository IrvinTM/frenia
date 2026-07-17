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
	"github.com/tiagomelo/go-clipboard/clipboard"
	"golang.org/x/term"
)

func Initial() {

	//Load the config
	configExists := util.CheckConfigFile()
	var pass *string

	config.GlobalConfig.Mode = "debug"

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

	} else {
		salt := crypt.GenerateRamdomSalt(32)
		configuration := types.Configuration{Salt: base64.StdEncoding.EncodeToString(salt)}
		util.UpdateConfig(&configuration)
	}

	if len(os.Args) > 1 {

		args := os.Args[1:]
		accName := args[1]

		if config.GlobalConfig.Mode == "debug" {
			fmt.Print("we got more than 1 arg")
			for i := range len(args) {
				fmt.Printf("\n arg %d : %s\n", i, args[i])
			}
		}

		switch args[0] {
		case "add":
				fmt.Println("case add...")
			if len(args) > 2 {
				accPass := args[2]
				if config.GlobalConfig.Mode == "debug" {
					fmt.Println("Adding a password")
				}
				fmt.Printf("adding password for %s", accPass)
				model.Save(*pass, accName, accPass)
				fmt.Println("Password saved")
				os.Exit(0)
			}

		case "get":
			if len(args) > 1 {

				if config.GlobalConfig.Mode == "debug" {

					fmt.Printf("Getting pass for %s\n", accName)
				}
				readPass, err := model.Read(*pass, accName)
				if err != nil {
					fmt.Println(err)
				}
				c := clipboard.New()
				if err := c.CopyText(readPass); err != nil {
					fmt.Printf("There was an error copying the pass to clipboard %s \n", err)
				}
				fmt.Printf("\n %s pass copied to clipboard", accName)
				//fmt.Println(readPass)
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

		fmt.Printf("\n The key being uses is: %d  %s \n", len(*pass), *pass)

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
