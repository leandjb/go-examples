package main

import "fmt"

type User struct {
	Name string
	Email string
}

func (self *User) setName(name string)  {
	self.Name = name
}


func main()  {
	user := User{ Name: "Lautaro", Email: "lautaro.martinez@inter.it"}
	fmt.Println("Before:", user)
	user.setName("Josef")
	fmt.Println("After:", user)



	
}