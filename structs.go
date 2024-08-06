package main

import "fmt"

type User struct {
	name  string
	email string
}

func main() {

	var cody User

	// cody.name = "Cody"
	// cody.email = "info@gmail.com"

	cody = User{name: "Cody", email: "info@gmail.com"}

	fmt.Println(cody.name, cody.email)
}
