package main

import (
	"fmt"
)

func main() {
	// dias := make(map[int]string)

	// dias[0] = "domingo"
	// dias[1] = "lunes"
	// dias[3] = "miercoles"

	usuarios := map[int]string{}

	usuarios[1] = "usuario #1"
	usuarios[2] = "usuario #3"

	for id, username := range usuarios {
		fmt.Println(id, username)

	}

}
