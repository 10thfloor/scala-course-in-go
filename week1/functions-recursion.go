package main

import "fmt"
import "math"

func sqrt(x float64) float64 {
	var sqrIter func(guess float64) float64

	sqrIter = func(guess float64) float64 {

		goodEnough := func(guess float64) bool {
			return math.Abs(guess*guess-x)/x < 0.001
		}

		improve := func(guess float64) float64 {
			return (guess + x/guess) / 2
		}

		if goodEnough(guess) {
			return guess
		}

		return sqrIter(improve(guess))
	}

	return sqrIter(1.0)
}

func main() {
	fmt.Println("Here comes some Maths!")
	fmt.Println("Square root of 40!", sqrt(40))
}
