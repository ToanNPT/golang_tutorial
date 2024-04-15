package main

import (
	"fmt"
	"time"
)

func main() {
	test(3)
}

func getDataSourceChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func test(numDBs int) {
	dbChan := getDataSourceChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}
