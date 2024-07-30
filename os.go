package main

import (
	"fmt"
	"runtime"
)

func main() {
	sistema := runtime.GOOS
	fmt.Printf("Estamos en %s\n", sistema)
}
