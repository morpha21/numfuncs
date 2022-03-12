package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if PrimeCheck(i) {
			fmt.Println(i)
		}
	}
}
