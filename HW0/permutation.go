package main

import "fmt"

func main() {
  fmt.Println(Permutation(1000, 2))
}

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
