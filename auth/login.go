package auth

import (
	"bufio"
	"fmt"
	"os"
)

func (a *Auth) Login() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Gagal Login")
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		var email, password string
		fmt.Println("\n===== Login =====")
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		found := false
		for _, u := range a.Users {
			if u.Email == email && u.Password == a.EncryptPassword(password) {
				fmt.Println("Login berhasil! Welcome", u.Firstname)
				fmt.Println("Press Enter to return to the main menu...")
				reader.ReadString('\n') 
				found = true
				a.Home(u.Firstname)
				break
			}
		}

		if found {
			break 
		} else {
			fmt.Println("Wrong Email or Password! Press Enter to try again...")
			reader.ReadString('\n') 
		}
	}
}
