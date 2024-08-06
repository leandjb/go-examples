package main

import "fmt"

type User struct {
	Name  string
	Email string
	Active bool
}

type Player struct {
	User User
	Id string
}

func (self *User) setName(name string) {
	self.Name = name
}

func (self *User) getName() string  {
	return self.Name
}

func main() {
	user := User{Name: "Lautaro", Email: "lautaro.martinez@inter.it"}
	fmt.Println("Before:", user)
	user.setName("Josef")
	fmt.Println("After:", user)
	

	fmt.Println("Other After:", user.getName())
}
