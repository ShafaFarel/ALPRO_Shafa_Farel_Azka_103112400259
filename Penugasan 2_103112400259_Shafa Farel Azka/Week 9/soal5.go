package main

import "fmt"

func main() {

	var gol1, gol2, gol3, gol4 int

	fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	max := gol1
	min := gol1

	if gol2 > max {
		max = gol2
	} else if gol2 < min {
		min = gol2
	}

	if gol3 > max {
		max = gol3
	} else if gol3 < min {
		min = gol3
	}

	if gol4 > max {
		max = gol4
	} else if gol4 < min {
		min = gol4
	}

	fmt.Println(max, min)

}
 //PSEUDOCODE

/*
program CariMaxMin
variabel
    gol1, gol2, gol3, gol4, max, min: integer
algoritma
    input(gol1, gol2, gol3, gol4)
    max = gol1
    min = gol1

    if gol2 > max then
        max = gol2
    else if gol2 < min then
        min = gol2
    end if

    if gol3 > max then
        max = gol3
    else if gol3 < min then
        min = gol3
    end if

    if gol4 > max then
        max = gol4
    else if gol4 < min then
        min = gol4
    end if

    output(max, min)
end program

*/