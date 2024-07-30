package main

import "fmt"

func esPar() {

	for i := 0; i <= 100; i++ {                            //for loop
		if i%2 == 0 {
			fmt.Println("Es par", i)
		}
	}
}

func bucleComoWhile()  {
	var num uint64 = 9999999999999999999
	count := 0

	for num > 0 {                                          // while loop
		num = num / 10
		count++
	}

	fmt.Println("La cantidad de digitos es:", count)
}

func bucleComoDoWhile() {

	var num = 1

	for ok:=true; ok; ok = num < 10 {                       //do while loop
		fmt.Println(num)
		num++
	}

}

func main() {
	//esPar()
	//bucleComoWhile()
	bucleComoDoWhile()
}
