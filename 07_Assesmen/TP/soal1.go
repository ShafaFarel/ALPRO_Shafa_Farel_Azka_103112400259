package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print(i * i)
		if i != n {
			fmt.Print(" ")
		}
	}
}
