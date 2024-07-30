package main

import "fmt"

func main()  {

	
	var calificacion int

	fmt.Println("Ingrese calificacion:")
	fmt.Scanf("%d", &calificacion)

	/*
	switch  {
	case calificacion == 10:
		fmt.Println("perfecto", calificacion)
	case calificacion == 8 || calificacion == 9:
		fmt.Println("bueno", calificacion)
	case calificacion == 6 || calificacion == 7:
		fmt.Println("raspando...", calificacion)
	case calificacion >= 0 && calificacion <= 5:
		fmt.Println("perdiste la materia", calificacion)
	default:
		fmt.Println("calificacion no valida")
	}
	*/

	switch calificacion {
	case 10:
		fmt.Println("perfecto", calificacion)
	case 8, 9:
		fmt.Println("bueno", calificacion)
	case 6, 7:
		fmt.Println("raspando...", calificacion)
	case 0,1,2,3,4,5:
		fmt.Println("perdiste la materia", calificacion)
	default:
		fmt.Println("calificacion no valida")
	}
}