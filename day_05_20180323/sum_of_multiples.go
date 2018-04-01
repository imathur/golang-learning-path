package main

import (
    "fmt"
)

func isMultiple(x int) bool {
    if (x % 3 == 0 || x % 5 == 0) {
         return true
    }
    return false
}

func getSum() int {
    sum := 0
    for i := 1 ; i < 1000 ; i++ {
	if (isMultiple(i)) {
	    sum += i
        }
    }
    return sum
}
func main() {
    sum := getSum()
    fmt.Println(sum)
}
