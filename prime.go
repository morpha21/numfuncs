package main

import "math"

func PrimeCheck(num int) bool {
	raiz := math.Sqrt(float64(num))
	for i := 2; i <= int(raiz); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true && num != 1
}
