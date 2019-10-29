package main

import "fmt"

func main() {
  // Test FactorialArray.
  fmt.Println(FactorialArray(10))
  // Test FibonacciArray
  fmt.Println(FibonacciArray(10))

  // Test MinArray
  // Create an integer array of 5 elements.
  integerArray := make([]int, 5)
  // Load the integer array.
  integerArray[0] = 11
  integerArray[1] = 21
  integerArray[2] = 6
  integerArray[3] = 20
  integerArray[4] = 33
  // Find min number. Fhould be 6.
  fmt.Println(MinArray(integerArray))

  // Test GCDArray.
  integerArrayTwo := make([]int, 3)
  integerArrayTwo[0] = 4
  integerArrayTwo[1] = 8
  integerArrayTwo[2] = 12
  // Should print 4 as GCD.
  fmt.Println(GCDArray(integerArrayTwo))
}

// Implement a function FactorialArray that takes an integer n and returns
// a slice of length n+1 whose k-th element is equal to k!
func FactorialArray(n int) []int {
  factorialArray := make([]int, n)
  // initialize first element of the array.
  factorialArray[0] = 1

  for i := 1; i < n; i++ {
    factorialArray[i] = factorialArray[i-1] * (i+1)
  }
  return factorialArray
}

// Write and implement a function FibonacciArray that takes an integer
// n as input and returns an array of length n whose k-th element is the
// k-th Fibonacci number.
func FibonacciArray(n int) []int {
  fibonacciArray := make([]int, n)

  // initialize first two elements of fibonacci.
  fibonacciArray[0] = 0
  fibonacciArray[1] = 1

  for i := 2; i < n; i++ {
    fibonacciArray[i] = fibonacciArray[i-2] + fibonacciArray[i-1]
  }
  return fibonacciArray
}

// Write and implement a function MinArray that takes an array of integers
// as input and returns the minimum of all these integers.
func MinArray(n []int) int {
  // initialize min number to the first array element.
  minNumber := n[0]
  for i := 1; i < len(n); i++ {
    if n[i] < minNumber {
      minNumber = n[i]
    }
  }
  return minNumber
}

// Write and implement a function GCDArray that takes an array of integers as
// input and generalizes the idea in TrivialGCD to return the greatest common
// divisor of all the integers in the array.
func GCDArray(intArray []int) int {
  gcd := 1
  minNumber := MinArray(intArray)
  for i := 1; i <= minNumber; i++ {
    for p := 0; p < len(intArray) ; p++ {
      if intArray[p] % i != 0 {
        gcd = 1
        break
      }
      gcd = i
    }
  }
  return gcd
}
