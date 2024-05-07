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
	fmt.Println("Testing for loops:")
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func testWhileLoop(age int) {
	fmt.Println("Testing while loops:")
	i := 1;
	for i < age {
		fmt.Println("Are you ", i, "?")
		i++;
	}	
	
}
