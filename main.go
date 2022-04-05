package main

import (
	"fmt"
	nc "numfuncs/numchecks"
	"time"
)

func main() {
	beginning := time.Now()
	counter := 0

	fmt.Println(beginning)

	go nc.PerfectCount(1, 4, &counter, 5, beginning)
	go nc.PerfectCount(2, 4, &counter, 5, beginning)
	go nc.PerfectCount(3, 4, &counter, 5, beginning)
	go nc.PerfectCount(4, 4, &counter, 5, beginning)

	//we are expecting to find 5 perfect numbers within 10minutes
	time.Sleep(10 * time.Minute)
}
