package main

import "fmt"

func main() {
    var batasAkhir int

    fmt.Print("Masukkan batas akhir: ")
    fmt.Scanln(&batasAkhir)

    fmt.Println("Bilangan genap dari 1 hingga", batasAkhir, ":")

    for i := 1; i <= batasAkhir; i ++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}     
    }
}