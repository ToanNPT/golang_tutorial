package main

import "fmt"

func main() {
	var array1 [5]string = [5]string{"Hello", "World", "Golang", "Programming", "Language"}
	for i := 0; i < len(array1); i++ {
		fmt.Printf("Index %d: %s\n", i, array1[i])
	}

	// Slicing an array
	subArray := array1[1:4]
	fmt.Println(subArray)

	subArray2 := make([]int, 3, 6)
	fmt.Println(subArray2)

	slice2 := []int{1, 10, 20}
	fmt.Println(slice2)

}
