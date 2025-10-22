package main

import (
	"fmt"
	"os"
)

func main() {
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
			fmt.Println("Menu Login dipilih")
		case 2:

			fmt.Println("Menu Login dipilih")
		case 3:

			fmt.Println("Menu Forgot Password dipilih")
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}

}
