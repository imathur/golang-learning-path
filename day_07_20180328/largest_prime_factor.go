package main

import (
	"math"
	"container/list"
	"fmt"
)

func isDivisible(number int, target int) bool {
    return (target >= 1 && number >= 1 && number % target == 0)
}

func isPrime(number int) bool {
	if (number < 2) {
		return false
	}
	for i := float64(2); i <= math.Sqrt(float64(number)); i++ {
		if (isDivisible(number, int(i))) {
			return false
		}
	}
	return true
}

func getLargestPrimeFactor(number int) int {
    var primeFactorsUnderSqrtNumber list.List
    //primeFactorsUnderSqrtNumber.PushBack(1)

    for i := float64(2); i <= math.Sqrt(float64(number)); i++ {
		if (isDivisible(number, int(i))) {
			factor := number / int(i)
			if (isPrime(factor)) {
				return factor
			}
			if (isPrime(int(i))) {
				primeFactorsUnderSqrtNumber.PushBack(int(i))
			}
		}
	}
	for factor := primeFactorsUnderSqrtNumber.Back(); factor != nil; factor.Prev() {
		if isPrime(factor.Value.(int)) {
			return factor.Value.(int)
		}
	}
	return number
}


func main() {
    number := 600851475143
    largestFactor := getLargestPrimeFactor(number)
    fmt.Println(largestFactor)
}
