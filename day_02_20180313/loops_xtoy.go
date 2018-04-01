package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    x, _ := strconv.Atoi(os.Args[1])
    y, _ := strconv.Atoi(os.Args[2])
    for i := x ; i < y ; i++ {
        fmt.Println(i)
    }
}
