package main

import "fmt"

type Animal interface {
	Comer()
	Dormir()
}

type Dog struct {
	Name string
}

func (self Dog) Comer() {
	fmt.Println("El perro come.")
}

func (self Dog) Dormir() {
	fmt.Println("El perro duerme.")
}

func doActions(animal Animal)  {
	animal.Comer()
	animal.Dormir()
}

func main()  {
	doggy := Dog{Name: "Lucas"}
	fmt.Println(doggy.Name)
	

	doActions(doggy)
}