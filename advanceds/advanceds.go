package main

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func aggregate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	println(aggregate(10, 20, add))
	println(aggregate(10, 20, sub))
}
