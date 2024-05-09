package main

import (
	"fmt" // for printing
	"time" // for accessing time
)

func main() {
	fmt.Println("Welcome to Program B: A Demonstration of Go!") // there is fmt.Println, fmt.Print, and fmt.Printf
	// testing functions/methods and output with all of these
	testForLoop() 
	testWhileLoop()
	testArray()
	testMath()
	testConditionals()
	testInput()
	fmt.Print("\nThank you for using my program.")
}

func testForLoop() {
	fmt.Println("\nTesting for loops")
	for i := 1; i < 10; i++ { // = operators are now :=
		for j := 0; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func testWhileLoop() {
	fmt.Println("\nTesting while loops")
	i := 1 // type is inferred
	var age int = 10 // type is specified, = instead of :=
	for i < age { // while loops use "for" keyword
		fmt.Printf("Are you %v?\n", i) // %v is value, %T is Go-syntax type of variable
		i++ // increments i by 1
	}
	fmt.Printf("You're %v!!!\n", age)
}

func testArray() {
	fmt.Println("\nTesting arrays") // arrays are printable
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
	fmt.Println("\nTesting math operations")
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
	fmt.Println("\nTesting conditionals")
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

func testInput() {
	fmt.Println("\nTesting input")
	fmt.Print("Enter your favorite number: ")
	var favNum int;
	fmt.Scan(&favNum) // &favNum is favNum's location in memory
	fmt.Print("Enter your two least favorite numbers: ")
	var badNum1 int // default value is 0
	var badNum2 int
	fmt.Scanf("%v %v", &badNum1, &badNum2) // separated by space
	fmt.Print("Enter your favorite word: ")
	var favWord string
	fmt.Scan(&favWord)
	fmt.Printf("Your favorite number is %v, your two least favorite numbers are %v and %v, and your favorite word is %v.\n", favNum, badNum1, badNum2, favWord) // printf works on multiple types of vars
}
