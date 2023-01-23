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

	//short variable declaration
	name := "toto"
	fmt.Println(name)

	//variable scope
	city := "Kingston"
	fmt.Println(city)
	{
		country := "Canada"
		fmt.Println(country)

	}
	//cant access country variable here

	//User input
	//var nom string
	//fmt.Println("Met ton nom:")
	//fmt.Scanf("%s", &nom)
	//fmt.Println("Le nom saisi est:", nom)

	//Multiple input
	var a string
	var b int
	fmt.Print("Add a string and a number")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count:", count)
	fmt.Println("error:", err)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

}
