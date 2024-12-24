package main

import "fmt"

// Fungsi utama
func main() {
    var num int
    fmt.Print("Masukkan sebuah bilangan: ")
    fmt.Scan(&num)

    sum := 0
    // Menjumlahkan faktor-faktor dari bilangan num
    for i := 1; i < num; i++ {
        if num%i == 0 {
            sum += i
        }
    }

    if sum == num {
        fmt.Println("Ya") // Bilangan sempurna
    } else {
        fmt.Println("Tidak") // Bukan bilangan sempurna
    }
}