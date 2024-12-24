package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var r float64
	var t float64

	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
	fmt.Print("Masukkan r: ")
	fmt.Scan(&r)

	fmt.Print("Masukkan t: ")
	fmt.Scan(&t)

	Volume := (1.0/3.0) * math.Pi * r * r * t
	fmt.Println("Volume kerucut adalah: ", Volume)
	}
}