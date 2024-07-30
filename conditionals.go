package main

import "fmt"

func main() {

	var calificacion int

	fmt.Println("Ingrese calificacion:")
	fmt.Scanf("%d", &calificacion)

	if calificacion == 10 {
		fmt.Print("Perfect!")
	} else if calificacion > 6 {
		fmt.Print("Raspado")
	} else {
		fmt.Print("Perdiste")
	}
	
}
