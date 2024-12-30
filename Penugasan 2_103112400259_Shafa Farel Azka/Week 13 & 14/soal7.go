package main

import "fmt"

func main() {
	var angka int
	var max, min, sum, count int
	var selesai bool

	max = -1000000
	min = 1000000

	for {
		fmt.Scan(&angka)
		if angka == 0 {
			if selesai {
				break
			}
			selesai = true
			continue
		}

		selesai = false
		if angka > max {
			max = angka
		}
		if angka < min {
			min = angka
		}
		sum += angka
		count++
	}

	if count > 0 {
		rata := float64(sum) / float64(count)
		fmt.Printf("Tertinggi: %d\n", max)
		fmt.Printf("Terendah: %d\n", min)
		fmt.Printf("Rata-rata: %.3f\n", rata)
	} else {
		fmt.Println("Tidak ada data yang valid.")
	}
}
//PSEUDOCODE

/*

program FindStatistics
variables
    angka, max, min, sum, count: integer
    rata: float
    selesai: boolean

algorithm
    max = -1000000
    min = 1000000
    selesai = false
    sum = 0
    count = 0

    while true do
        input(angka)
        
        if angka == 0 then
            if selesai then
                break
            end if
            selesai = true
            continue
        end if

        selesai = false
        if angka > max then
            max = angka
        end if
        if angka < min then
            min = angka
        end if

        sum = sum + angka
        count = count + 1
    end while

    if count > 0 then
        rata = sum / count
        output("Highest: ", max)
        output("Lowest: ", min)
        output("Average: ", rata)
    else
        output("No valid data.")
    end if
end program
*/