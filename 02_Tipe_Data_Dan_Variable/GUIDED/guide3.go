package main

import (
	"fmt"
)

func main() {

	var x float32

	fmt.Print("Masukan bilangan x: ")
	fmt.Scan(&x)

	var opearasi float32 = 2/(x+5) + 5 //Rumus dari fungsi f(x)
	fmt.Println("Hasil dari:", x, "adalah", opearasi)
}