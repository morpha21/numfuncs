package numchecks

import "math"

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
