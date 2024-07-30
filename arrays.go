package main

import "fmt"

func main() {
	//var numeros[5] int
	//numeros := [5]int{100, c00, 300, 400, 500}
	//numeros := [...]int{100, 200, 300, 400, 500}
	slice := make([]int, 0, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
