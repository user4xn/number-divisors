package main

import (
	"fmt"
)

func main() {
	n := 36
	for i := 1; i <= n; i++ {
		divisorsChallange(i)
	}
}

func divisorsChallange(num int) {
	divs := []int{}     //DIVISORS NEST
	evenDivs := []int{} //EVEN DIVS NEST

	// INSTEAD OF CHECKING EVERY NUMBER ~ NUM, ONLY CHECK NUMBER ~ (ROOT OF NUM)
	// AVOIDING USING SQRT LIBRARY, MUCH FASTER I THINK. CZ LIB STILL NEED TO PROCESSING
	for i := 1; i*i <= num; i++ {
		// EX (36%2 == 0) == 36 WAS DIVISORS, DIVIDE NUM EVENLY = DIVISORS
		if num%i == 0 {
			divs = append(divs, i) // APPEND TO NEST

			// CHECK IF THE DIVS 2 WAS EVEN
			if i%2 == 0 {
				evenDivs = append(evenDivs, i)
			}

			// EX (36/2 = 18 , 18 != 2) == 36/2 WAS DIVISORS
			if num/i != i {
				divs = append(divs, num/i) // APPEND TO NEST

				// CHECK IF THE DIVS 2 WAS EVEN
				if (num/i)%2 == 0 {
					evenDivs = append(evenDivs, num/i)
				}
			}
		}
	}

	// PRINT THE SUMMARY
	title := fmt.Sprintf("======================%d=====================", num)
	log := fmt.Sprintf("%d are divisors of %d", divs, num)
	logTotal := fmt.Sprintf("got total %d divisors", len(divs))
	logEven := fmt.Sprintf("%d are the special number", evenDivs)
	logTotalEven := fmt.Sprintf("there is %d special number", len(evenDivs))

	fmt.Println(title)
	fmt.Println(log)
	fmt.Println(logTotal)
	fmt.Println(logEven)
	fmt.Println(logTotalEven)
}
