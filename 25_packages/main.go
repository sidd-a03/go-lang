package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sidd-a03/go-lang/auth"
	"github.com/sidd-a03/go-lang/user"
)

func main() {
	auth.LoginWithCredentials("sidd", "123123")

	session := auth.GetSession()
	fmt.Println(session)

	newUser := user.User{
		Email:    "user@example.com",
		Name:     "John",
		Password: "123123",
	}

	color.Red(newUser.Email)
}
