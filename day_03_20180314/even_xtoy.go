package main

import (
    "fmt"
    "strconv"
    "os"
)

func main() {
    x, _ := strconv.Atoi(os.Args[1])
    y, _ := strconv.Atoi(os.Args[2])
    for i := x ; i < y; i++ {
        if i % 2 == 0 {
            fmt.Println(i)
	}
    }
}
