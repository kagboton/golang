package main

import (
	"fmt"
	"reflect"
	"strconv"
)

const PI float64 = 3.14 //global constant

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

	//Type of a variable with %T
	var grades int = 42
	fmt.Printf("variable grades = %v is of type %T", grades, grades)

	//Type of a variable with TypeOf function from reflect package
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))

	//Type casting - convert types
	var i int = 12
	var f float64 = float64(i) //convert int to float
	fmt.Println(f)

	var s = strconv.Itoa(i) //convert int to string
	fmt.Printf("%q", s)

	//Constants - type is optional
	const myConst string = "my string constant" //typed constant
	const ageee = 16                            //untyped constant
	const t_name, t_age = "Toto", 15            //untyped constant
}
