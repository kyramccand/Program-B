package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Program B: A Demonstration of Go!")
	testForLoop()
	testWhileLoop()
	testArray()
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

func testWhileLoop() {
	fmt.Println("\nTesting while loops:")
	i := 1
	age := 10
	for i < age { // while loops use for keyword
		fmt.Printf("Are you %v?\n", i) // printf
		i++
	}
	fmt.Printf("You're %v!!!\n", age)
}

func testArray() {
	fmt.Println("\nTesting arrays:") // arrays are printable in Go!
	var dayArray = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} // length is inferred
	monthArray := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"} // length is given
	fmt.Println("Days of the week:", dayArray)
	fmt.Println("Months of the year:", monthArray)
	weekday := int(time.Now().Weekday()) // type is inferred
	month := int(time.Now().Month()) - 1;
	var num int = 10; // type is given
	fmt.Println("In", num, "days, it will be a", dayArray[(weekday + num) % 7] + ".\nIn", num, "months, it will be", monthArray[(month + num) % 12] + ".")
}
