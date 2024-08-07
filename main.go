package main

import (
	"./PackageGo"
	"fmt"
)

func main() {
	curso := PackageGo.Curso{Titulo: "Curso profesional de GO"}
	fmt.Println(curso)
}
