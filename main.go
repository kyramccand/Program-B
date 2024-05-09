package main

import (
	"fmt" // for printing
	"time" // for accessing time
)

func main() {
	fmt.Println("Welcome to Program B: A Demonstration of Go!") // there is fmt.Println, fmt.Print, and fmt.Printf
	testForLoop() // testing functions/methods with all of these
	testWhileLoop()
	testArray()
	testMath()
	testConditionals()
}


func testForLoop() {
	fmt.Println("\nTesting for loops:")
	for i := 1; i < 10; i++ { // = operators are now :=
		for j := 0; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func testWhileLoop() {
	fmt.Println("\nTesting while loops:")
	i := 1 // type is inferred
	var age int = 10 // type is specified, = instead of :=
	for i < age { // while loops use "for" keyword
		fmt.Printf("Are you %v?\n", i) // %v is value, %T is Go-syntax type of variable
		i++ // increments i by 1
	}
	fmt.Printf("You're %v!!!\n", age)
}

func testArray() {
	fmt.Println("\nTesting arrays:") // arrays are printable in Go!
	var dayArray = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} // length is inferred
	monthArray := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"} // length is given
	fmt.Println("Days of the week:", dayArray)
	fmt.Println("Months of the year:", monthArray)
	weekday := int(time.Now().Weekday())
	month := int(time.Now().Month()) - 1;
	var num int = 10;
	fmt.Println("In", num, "days, it will be a", dayArray[(weekday + num) % 7] + ".\nIn", num, "months, it will be", monthArray[(month + num) % 12] + ".")
}

func testMath() {
	fmt.Println("\nTesting math operations:")
	a := 10 // Go is strict about unused variables
	b := 5
	fmt.Println(a, "+", b, "=", a+b)
	fmt.Println(a, "-", b, "=", a-b)
	fmt.Println(a, "*", b, "=", a*b)
	fmt.Println(a, "/", b, "=", a/b)
	fmt.Println(a, "%", b, "=", a%b)
	c:= 10.5
	d := 5.5
	// floats print as ints if they end in .0
	fmt.Println(c, "+", d, "=", c+d)
	fmt.Println(c, "-", d, "=", c-d)
	fmt.Println(c, "*", d, "=", c*d)
	fmt.Println(c, "/", d, "=", c/d)
	// % is not defined for floats
	// +, -, *, and / are not defined for variables of different types
	// to perform operations between an int and float, you must first 
	// cast one type to the other
}

func testConditionals() {
	fmt.Println("\nTesting conditionals:")
	a := 0
	b := 1
	c := 5
	if a < b {
		fmt.Println(a, "is less than", b)
	} else if a == b { // must be the same line as closing bracket
		fmt.Println(a, "is equal to", b)
	} else {
		fmt.Println(a, "is greater than", b)
	}
	if a == c || b == c {
		fmt.Println(a, "or", b, "is equal to", c)
		if a == c && b == c { // nested if statement
			fmt.Println(a, "and", b, "are equal to", c)
		}
	} else {
		fmt.Println(a, "and", b, "are not equal to", c)
	}
	
}