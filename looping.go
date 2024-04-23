package main

import "fmt"

func main() {
	var num int

	fmt.Println("Mini Challenge 1: FizzBuzz App | Dea Ananda Gunawan")
	fmt.Print("Input a number: ")
	fmt.Scan(&num)
	fmt.Println("The number you pick: ", num)
	
	fmt.Println("=========================")

	for i := 1; i <= num; i++ {
		// fmt.Println("Num = ", i);
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i % 5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}