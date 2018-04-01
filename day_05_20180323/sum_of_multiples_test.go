package main

import (
    "testing"
)

func TestIfDivisibleBy3Or5(t *testing.T) {
    x1 := 10 // divisible by 5
    if (isMultiple(x1) != true) {
         t.Fail()
    }
    x2 := 22 // divisible by neither
    if (isMultiple(x2) != false) {
         t.Fail()
    }
    x3 := 30 // divisible by both
    if (isMultiple(x3) != true) {
         t.Fail()
    }
}

func TestIfSumIsCorrect(t *testing.T) {
    if (getSum() != 233168) {
        t.Fail()
    }
}
