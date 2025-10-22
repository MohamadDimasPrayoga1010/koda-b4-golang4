package auth

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func (a *Auth) EncryptPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (a *Auth) Register() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Gagal register:", r)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for { 
		var firstName, lastName, email, password string

		for { 
			var confirmPassword string
			fmt.Println("=====Register=====")
			fmt.Print("Whats is your first name: ")
			fmt.Scan(&firstName)
			fmt.Print("Whats is your last name: ")
			fmt.Scan(&lastName)
			fmt.Print("Whats is your Email: ")
			fmt.Scan(&email)
			fmt.Print("Enter a strong password: ")
			fmt.Scan(&password)
			fmt.Print("Confirm your Password: ")
			fmt.Scan(&confirmPassword)

			if password != confirmPassword {
				fmt.Println("Wrong confirm password, press Enter to try again...")
				reader.ReadString('\n') 
				continue
			}
			break
		}

		fmt.Println("\nIs it correct?")
		fmt.Println("FirstName:", firstName)
		fmt.Println("LastName :", lastName)
		fmt.Println("Email    :", email)

		var confirm string
		fmt.Print("Continue? (y/n): ")
		fmt.Scan(&confirm)
		confirm = strings.ToLower(confirm)

		if confirm == "y" {
			a.Users = append(a.Users, &User{
				Firstname: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  a.EncryptPassword(password),
			})

			fmt.Println("Register success, press enter to back..")
			reader.ReadString('\n') 
			break 
		} else if confirm == "n" {
			fmt.Println("Registration canceled, please refill the data...\n")
			continue 
		} else {
			fmt.Println("Invalid input, please try again...\n")
			continue
		}
	}
}
