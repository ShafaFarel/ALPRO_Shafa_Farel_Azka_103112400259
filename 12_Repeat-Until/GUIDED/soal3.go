package main

import "fmt"

func main() {

	var x,y int

	fmt.Print("Masukkan bilangan pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan kedua (y): ")
	fmt.Scan(&y)

	for selesai := false; !selesai;{
		x = x-y

		fmt.Println(x)
		selesai = x <= 0
	}

	fmt.Println(x==0)
}