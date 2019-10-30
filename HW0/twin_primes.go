package main

import "fmt"

func main() {
  fmt.Println(NextTwinPrimes(2))
  fmt.Println(NextTwinPrimes(41))
  fmt.Println(NextTwinPrimes(191))
}

// Returns true if n is prime.
func IsPrime(n int) bool {
  // Return false if n is 1.
  if n == 1 {
    return false
  }

  for p := 2; p < n; p++ {
    if (n % p == 0) {
      return false
    }
  }
  return true
}

// Write and implement a function NextTwinPrimes that takes an integer
// n as input and returns the smallest pair of twin primes that are both
// larger than n, using IsPrime as a subroutine.
func NextTwinPrimes(n int) (a, b int) {

  // Initialize primeOne to n.
  primeOne := n

   // Increment primeOne to a number larger than n.
  primeOne += 1

  // Increment primeOne until it becomes a prime number.
  for IsPrime(primeOne) == false {
    primeOne += 1
  }

  // Initiate primeTwo to a number larger than primeOne.
  primeTwo := primeOne + 1

  // Increment primeTwo until it becomes a prime number.
  for IsPrime(primeTwo) == false {
    primeTwo += 1
  }

  return primeOne, primeTwo
}
