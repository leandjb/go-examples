package main

import (
	"fmt"
)

func main()  {

	var name string
	var age int
	var height float32

	
	fmt.Println("Ingrese primer nombre:")
	fmt.Scanf("%s", &name)

	fmt.Println("Ingrese edad:")
	fmt.Scanf("%d", &age)

	fmt.Println("Ingrese altura (cm):")
	fmt.Scanf("%f", &height)

	fmt.Printf("Su nombre es %s\n", name)
	fmt.Printf("Su edad es %d\n", age)
	fmt.Printf("Su altura es %.2f\n", height)
	

}