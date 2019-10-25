package main

import "fmt"

func main() {
  fmt.Println(Combination(1000, 2))
}

// the permutation function will be used to design the combination function.
func Permutation(n, k int) int {
  // handle input when n < k
  if n < k {
    panic("Error: input n must be greater than input k.")
  }

  answer := 1
  for i := n; i > (n - k); i-- {
    answer = answer * i
  }
  return answer
}

func Combination(n, k int) int {
  // use combination symmetry to avoid large numbers, if possible.
  if (n - k) < k {
    k = n - k
  }

  // calculate the value of P(n, k).
  permutationVal := Permutation(n, k)

  kFactorial := 1
  for k > 1 {
    kFactorial *= kFactorial * k
    k--
  }

  // C(n,k) = P(n,k)/k!
  combinationVal := permutationVal / kFactorial
  return combinationVal
}
