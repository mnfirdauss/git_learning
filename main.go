package main

import "fmt"

func main() {
	fmt.Println("ini baru")

	// calling fullName function
	fullName("Berryl", "Radian")
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int32) int32 {
	return a + b
}

// make fullName function
func fullName(firstName, lastName string) {
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
}
