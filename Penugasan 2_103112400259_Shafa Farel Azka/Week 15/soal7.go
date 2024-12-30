package main

import "fmt"

func main() {

	var x, j, bilangan, nX, nZero int

	fmt.Scan(&x)

	for j = 1; j <= 9; j++ {
		fmt.Scan(&bilangan)

		if bilangan == x {
			nX++
		} else {
			nZero++
		}
	}

	if nX > nZero {
		fmt.Print("Modus = ", x)
	} else {
		fmt.Print("Modus = ", nZero)
	}

}

//PSEUDO CODE
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