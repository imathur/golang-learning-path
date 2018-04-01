package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    n, _ := strconv.Atoi(os.Args[1])

    for i := 1 ; i < n ; i++ {
	div_by_3 := (i % 3 == 0)
	div_by_5 := (i % 5 == 0)

	if (div_by_3 || div_by_5) {
	    if div_by_3 {
		fmt.Print("Fizz ")
	    }
	    if div_by_5 {
		fmt.Print("Buzz")
	    }
	    fmt.Println()
        } else {
            fmt.Println(i)
        }
    }
}
