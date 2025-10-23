package auth

import (
	"bufio"
	"fmt"
	"os"
)

type MenuLoggedIn struct {
	name string
	menu []string
}
type MenuDisplay interface {
	OutputMenu()
}

func (m MenuLoggedIn) OutputMenu() {
	fmt.Printf("\n===== HOME =====\nWelcome %s!\n", m.name)
	for i, item := range m.menu {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("\n0. Exit")
	fmt.Print("Choose menu: ")
}

func (a *Auth) Home(name string) {
	reader := bufio.NewReader(os.Stdin)

	var menuDisplay MenuDisplay = MenuLoggedIn{
		name: name,
		menu: []string{"List all users", "Logout"},
	}

	for {
		menuDisplay.OutputMenu()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\n--- List all users ---")
			for i, u := range a.Users {
				fmt.Printf("%d.\n", i+1)
				fmt.Printf("Full Name: %s %s\n", u.Firstname, u.LastName)
				fmt.Printf("Email: %s\n", u.Email)
				fmt.Printf("Password: %s\n\n", u.Password)
			}
			fmt.Print("Press Enter to back...")
			reader.ReadString('\n')

		case 2:
			fmt.Println("Logging out...")
			reader.ReadString('\n')
			return

		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
			return

		default:
			fmt.Println("Invalid choice")
			fmt.Println("Press Enter to return to Home...")
			reader.ReadString('\n')
		}
	}
}
