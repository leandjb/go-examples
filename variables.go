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

func main()  {
	printVariableSimple()
	printConstants()
}
