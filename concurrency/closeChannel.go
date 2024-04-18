package main

import (
	"fmt"
	"time"
)

func testFunc(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func countReports(numSenCh chan int) int {
	totalReports := 0
	for num, ok := <-numSenCh; ok; num, ok = <-numSenCh {
		totalReports += num
	}
	return totalReports
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func concurrentFibonacci(n int) {
	ch := make(chan int, 5)
	go fibonacci(n, ch)

	for val := range ch {
		fmt.Println(val)
	}
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	testFunc(3)
	fmt.Println("========================")
	concurrentFibonacci(10)
}
