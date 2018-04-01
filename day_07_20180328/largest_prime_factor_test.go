package main

import (
    "testing"
)

func TestShouldReturnTrueIfNumberDivisibleByTarget(t *testing.T) {
    if (isDivisible(16, 4) != true) {
        t.Fail()
    }
}

func TestShouldReturnFalseIfNumberNotDivisibleByTarget(t *testing.T) {
    if (isDivisible(16, 3) != false) {
        t.Fail()
    }
}

func TestShouldReturnFalseIfDivisibilityTargetIsNotPositive(t *testing.T) {
   if (isDivisible(16, -4) != false) {
       t.Fail()
   }
}

func TestShouldReturnFalseIfDivisibilityNumberIsNotPositive(t *testing.T) {
   if (isDivisible(-16, 4) != false) {
       t.Fail()
   }
}

func TestShouldReturnTrueIfNumberIsPrime(t *testing.T) {
    if (isPrime(5) != true) {
        t.Fail()
    }
}

func TestShouldReturnFalseIfNumberIsNotPrime(t *testing.T) {
    if (isPrime(6) != false) {
        t.Fail()
    }
}

func TestShouldReturnFalseIfNumberIsLessThan2(t *testing.T) {
    if (isPrime(-5) != false) {
        t.Fail()
    }
}

func TestShouldReturnLargestPrimeFactorWhenNumberGreaterThanItsSqrt (t *testing.T) {
    if (getLargestPrimeFactor(74) != 37) {
        t.Fail()
    }
}

func TestShouldReturnLargestPrimeFactorWhenNumberLessThanItsSqrt (t *testing.T) {
    if (getLargestPrimeFactor(36) != 3) {
        t.Fail()
    }
}

func TestShouldReturnNumberWhenNumberIsPrime (t *testing.T) {
    if (getLargestPrimeFactor(37) != 37) {
        t.Fail()
    }
}
