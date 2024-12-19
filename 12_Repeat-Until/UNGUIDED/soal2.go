package main

import "fmt"

func main() {
    var desimal float64

    fmt.Print("Masukkan bilangan desimal: ")
    fmt.Scan(&desimal)

    pembulatan := int(desimal)
    if desimal > float64(pembulatan) {
        pembulatan++
    }

    sum := desimal

    // Lakukan penjumlahan hingga mencapai pembulatan ke atas
    for {
        if int(sum) >= pembulatan {
            // Jika sum sudah mencapai atau melebihi upperBound, tampilkan sebagai int
            fmt.Println(int(sum))
        } else {
            // Tampilkan sum sebagai float dengan satu angka di belakang koma
            fmt.Printf("%.1f\n", sum)
        }

        // Jika sudah mencapai atau melebihi upperBound, keluar dari loop
        if int(sum) >= pembulatan {
            break
        }

        // Tambahkan 0.1 ke sum untuk iterasi berikutnya
        sum += 0.1
    }
}