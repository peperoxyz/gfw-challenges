package main

import "fmt"

func sesiSatu() {
	fmt.Println("Hello World")

	// variable with data type
	var title string = "Golang For Women"
	var totalMember int = 13

	fmt.Println("Course title is:", title)
	fmt.Println("Total member is:", totalMember)

	// initiate first, assign later
	var title2 string
	title2 = "Golang For Men"
	var totalMember2 int
	totalMember2 = 20

	fmt.Println("Course title 2 is:", title2)
	fmt.Println("Total Member 2 is:", totalMember2)

	// variable without data type: Golang can analize what is the data type
	// short declaration type
	name:= "Dea Ananda Gunawan"
	age:= 22

	fmt.Println("My name is", name, "I am", age, "years old.")

	// multiple declaration variable
	// underscore variable

	// ====================================================

	// Data type in Go
	// Integer
	first := 22
	second := -12

	fmt.Printf("First data value is: %d\n", first)
	fmt.Printf("Second data type is: %T\n", second)
}