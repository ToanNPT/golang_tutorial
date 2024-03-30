package main

import (
	"fmt"
)

type DividedByZeroError struct {
	divisor float64
}

func (e DividedByZeroError) Error() string {
	return fmt.Sprintf("We can not perform operator. Divided by zero error: divisor = %.2f", e.divisor)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, DividedByZeroError{divisor: divisor}
	} else {
		return dividend / divisor, nil
	}
}

func main() {
	rs, err := divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Result: %.2f", rs)
	}
}
