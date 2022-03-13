package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if PrimeCheck(i) {
			fmt.Println(i)
		}
	}

	for i := 1; i < 8200; i++ {
		if PerfectCheck(i) {
			fmt.Println(i)
		}
	}
}
