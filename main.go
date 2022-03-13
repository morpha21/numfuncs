package main

import (
	"fmt"
	nc "numfuncs/numchecks"
)

func main() {
	fmt.Println("Some prime numbers:")
	for i := 1; i < 20; i++ {
		if nc.PrimeCheck(i) {
			fmt.Println(i)
		}
	}

	fmt.Println()

	fmt.Println("some perfect numbers:")
	for i := 1; i < 8200; i++ {
		if nc.PerfectCheck(i) {
			fmt.Println(i)
		}
	}
}
