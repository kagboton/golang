package main

import "fmt"

func main() {
	var hello string = "Hello"

	var firstname string = "Harry"
	var lastname string = "Potter"

	var grade int = 12

	fmt.Println(hello, firstname, lastname)

	//print format %v
	fmt.Printf("Nice to meet you %v", firstname)

	fmt.Print("\n")

	//print format %d
	fmt.Printf("Welcome to grade %d", grade)
}
