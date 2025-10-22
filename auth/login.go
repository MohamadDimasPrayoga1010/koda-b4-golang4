package auth

import (
	"fmt"
)

func (a *Auth) Login() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Gagal Login")
		}
	}()
	// reader := bufio.NewReader(os.Stdin)

	var email, password string
	fmt.Println("\n=====Login=====")
	fmt.Print("Enter your email : ")
	fmt.Scan(&email)
	fmt.Print("Enter your password")
	fmt.Scan(&password)

	for _, u := range a.Users {
		if u.Email == email && u.Password == a.EncryptPassword(password) {
			fmt.Println("Login berhasil! Welcome", u.Firstname)
			return
		}
	}

	fmt.Println("Email atau Password salah!")

}
