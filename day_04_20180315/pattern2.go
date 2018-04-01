package main

import "fmt"

func main() {
    for i := 1 ; i <= 5 ; i++ {
	for j := 5 ; j >=i ; j-- {
	    fmt.Print("* ")
        }
        fmt.Println()
    }
}
