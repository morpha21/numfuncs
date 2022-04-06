package numchecks

import (
	"fmt"
	"math"
	"time"
)

//checks whether num is a perfect number or not
func PerfectCheck(num int) bool {
	div := 1
	raiz := math.Sqrt(float64(num))

	for i := 2; i <= int(raiz); i++ {
		if num%i == 0 {
			div += i + num/i
		}
	}
	return num == div && num != 1
}

//counts perfect numbers starting from <num>, and then jumping <jump> steps, until it prints <quantity> numbers
//the jump's purpose is to call the function concurrently, from different starting points
func PerfectCount(num int, jump int, count *int, quantity int, beginning time.Time) {
	for *count < quantity {
		if PerfectCheck(num) {
			fmt.Printf("Perfect number:\t\t%v\ntime since beginning:\t%v\n\n", num, time.Since(beginning))
			*count++
		}
		num += jump
	}

}

//checks whether num is a prime number or not
func PrimeCheck(num int) bool {
	raiz := math.Sqrt(float64(num))
	for i := 2; i <= int(raiz); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true && num != 1 //1 is not a prime number
}
