package main

import (
    "fmt"
)

func isNumberGreaterThanTarget(number int, target int) bool {
    return (number > target)
}

func isEven(x int) bool {
    return (x % 2 == 0 && x > 0)
}

func getSumOfEvenFibonacci(highestValueItem int) int {
    first := 1
    second := 2
    third := 0
    sum := 2
    for {
	third = first + second
	if isNumberGreaterThanTarget(third, highestValueItem) {
            break
	}
	if (isEven(third)) {
	    sum += third
        }
	first = second
	second = third
    }
    return sum
}

func main() {
    highestValueItem := 4000000
    sum := getSumOfEvenFibonacci(highestValueItem)
    fmt.Println(sum)
}
