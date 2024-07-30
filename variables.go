package main

import "fmt"

const League string = "Serie A Enilive"

func printVariableSimple() {

	team1 := "INTER"
	team2 := "JUVENTUS"
	team3 := "MILAN"

	fmt.Println(team1, team3, team2)
}

func printConstants()  {
	fmt.Println(League)
}

func imprimirSemana()  {
	const (
		domingo int = iota + 1
		lunes
		martes
		miercoles
		jueves
		viernes
		sabado
	)

	fmt.Println(domingo)
	fmt.Println(lunes)
	fmt.Println(martes)
	fmt.Println(miercoles)
	fmt.Println(jueves)
	fmt.Println(viernes)
	fmt.Println(sabado)
}

func main()  {
	printVariableSimple()
	printConstants()
	imprimirSemana()
}
