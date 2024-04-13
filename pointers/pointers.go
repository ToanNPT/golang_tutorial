package main

import "fmt"

func main() {
	var a int = 10
	z := &a
	fmt.Printf("Value of z: %d\n", *z)
}

func increase(a int) int {
	return a + 1
}

func inscrease(a *int) {
	*a = *a + 1
}
