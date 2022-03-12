package main

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
