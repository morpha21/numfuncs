package main

import (
	nc "numfuncs/numchecks"
	"os"
	"strconv"
	"time"
)

func main() {
	beginning := time.Now()
	counter := 0
	channel := make(chan bool)

	numOfPerfs, _ := strconv.Atoi(os.Args[1])
	goNum, _ := strconv.Atoi(os.Args[2])

	for i := 1; i <= goNum; i++ {
		go nc.PerfectCount(i, goNum, &counter, numOfPerfs, beginning, channel)
	}

	<-channel
}
