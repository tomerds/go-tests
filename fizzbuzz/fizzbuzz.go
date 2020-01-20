package main

import "fmt"

// FizzBuzz calculator
func FizzBuzz(i int) string {
	result := ""
	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	} else {
		result = fmt.Sprintf("%d", i)
	}

	return result
}

// Iterator loops fizzbuzz over a given amount
func Iterator(amount int) string {
	totalOutput := ""
	for i := 1; i <= amount; i++ {
		output := FizzBuzz(i)
		totalOutput += output + "\n"
	}
	return totalOutput
}

func main() {
	fmt.Println(Iterator(100))
}
