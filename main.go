package main

import "log"

// Addition is a variadic function, i.e  accepts any numbers of arguments
// sums all the args and return the sum
func Addition(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// Subtraction 2 numbers and returns subtraction result
func Subtraction(x, y int) int {
	return x - y
}

func Division(x, y int) int {
	result := x / y
	log.Println("result", result)
	return result
}

// Multiplication is variadic function accepts array of numbers and retuylr muitiplication result
func Multiplication(numbers ...int) int {
	// 0th index is assigned as result to be multiplied with other elements, note range[1:] below
	result := numbers[0]
	for _, v := range numbers[1:] {
		result *= v
	}

	return result
}

func SomeFunc() {
	log.Println("no test coverage")
}

func main() {

}
