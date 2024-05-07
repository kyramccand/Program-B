package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Program B: A Demonstration of Go!")
	testForLoop()
	testWhileLoop(10)
}


func testForLoop() {
	fmt.Println("\nTesting for loops:")
	for i := 1; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func testWhileLoop(age int) {
	fmt.Println("\nTesting while loops:")
	i := 1;
	for i < age { // while loops use for keyword
		fmt.Printf("Are you %v? ", i) // printf
		i++;
	}
	fmt.Printf("You're %v!!!\n", age)
}
