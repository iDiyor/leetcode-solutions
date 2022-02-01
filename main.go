package main

import "fmt"

func counter(inChan chan int) {
	for i := 0; i < 10; i++ {
		inChan <- i
	}
	close(inChan)
}

func squarer(inChan chan int, outChan chan int) {
	for val := range inChan {
		outChan <- val * val
	}
	close(outChan)
}

func printer(inChan chan int) {
	for val := range inChan {
		fmt.Printf("Value: %d\n", val)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go counter(chan1)
	go squarer(chan1, chan2)
	printer(chan2)
}
