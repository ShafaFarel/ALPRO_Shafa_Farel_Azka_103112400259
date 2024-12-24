package main

import "fmt"

func main () {

	var a int
	var b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i)
	}
}