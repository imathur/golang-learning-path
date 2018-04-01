package main

import (
    "testing"
)

func TestShouldReturnTrueIfNumberGreaterThanTarget(t *testing.T) {
    if (isNumberGreaterThanTarget(16, 10) != true) {
        t.Fail()
    }
}

func TestShouldReturnFalseIfNumberEqualToTarget(t *testing.T) {
    if (isNumberGreaterThanTarget(10, 10) != false) {
        t.Fail()
    }
}

func TestShouldReturnFalseIfNumberLessThanTarget(t *testing.T) {
    if (isNumberGreaterThanTarget(6, 10) != false) {
        t.Fail()
    }
}

func  TestShouldReturnTrueIfNumberIsEven(t *testing.T) {
    if (isEven(4) != true) {
        t.Fail()
    }
}

func  TestShouldReturnFalseIfNumberIsOdd(t *testing.T) {
    if (isEven(3) != false) {
        t.Fail()
    }
}

func  TestShouldReturnFalseIfNumberIsNegative(t *testing.T) {
    if (isEven(-5) != false) {
        t.Fail()
    }
}

func  TestShouldReturnTrueIfSumOfEvenFibonacciNumbersIsCorrect(t *testing.T) {
    if (getSumOfEvenFibonacci(100) != 44) {
	t.Errorf("\nExpected: 44\nActual: %d", getSumOfEvenFibonacci(100))
    }
}


