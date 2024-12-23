package main

import (
	"fmt"
	"myapp/24_Packages/auth"
	"myapp/24_Packages/user"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("Pikachu", "pikapika")
	session := auth.GetSession()
	fmt.Println("Session:", session)

	user := user.User{
		Email:    "pikachu@pokemon.com",
		Password: "pikapika",
	}

	fmt.Println("User:", user)

	color.Red(user.Email)
	color.Green(user.Password)
}
