package auth

import (
	"bufio"
	"fmt"
	"os"
)

func (a *Auth) ForgotPassword() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Forgot Password ---")
	fmt.Print("Enter your email: ")
	var email string
	fmt.Scan(&email)

	var userFound *User
	for _, u := range a.Users {
		if u.Email == email {
			userFound = u
			break
		}
	}

	if userFound == nil {
		fmt.Println("Email not found! Press Enter to return...")
		reader.ReadString('\n')
		return
	}

	for {
		var password, confirmPassword string
		fmt.Print("Enter a strong password: ")
		fmt.Scan(&password)
		fmt.Print("Confirm your password: ")
		fmt.Scan(&confirmPassword)

		if password != confirmPassword {
			fmt.Println("Wrong confirm password, press enter to back!")
			reader.ReadString('\n') 
			continue
		}

		userFound.Password = a.EncryptPassword(password)
		fmt.Println("Password changed successfully! Press Enter to return...")
		reader.ReadString('\n')
		break
	}
}
