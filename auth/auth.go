package auth


type Auth struct {
	Users []*User
}

type Authenticator interface {
	Register()
	Login()
}

