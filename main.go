package main

import (
	"fmt"
	"os"

	"main.go/auth"
)

func main() {
	authSystem := &auth.Auth{Users: []*auth.User{}}
	var choice int
	for {
		fmt.Println("\n===== WELCOME TO SYSTEM =====")
		fmt.Println("\n1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Forgot Password")
		fmt.Println("0. Exit")
		fmt.Print("\nChoose a menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			authSystem.Register()
		case 2:
			authSystem.Login()
		case 3:
			authSystem.ForgotPassword()
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}

}
